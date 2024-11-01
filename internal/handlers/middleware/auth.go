package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strings"
	"text-to-api/internal/domain"
	hdlPorts "text-to-api/internal/handlers/ports"
	"text-to-api/internal/ports"
)

// AuthMdlwHdl is the receiver for the authentication middleware. It holds the
// required components to perform authentication.
type AuthMdlwHdl struct {
	authService ports.AuthService
	logger      ports.Logger
	reqCtxHdl   hdlPorts.RequestContextHandler
}

// NewAuthMdlwHdl returns a new instance of AuthMdlwHdl. It returns an error if any of the
// required components are nil.
func NewAuthMdlwHdl(authService ports.AuthService, logger ports.Logger, reqCtxHdl hdlPorts.RequestContextHandler) (*AuthMdlwHdl, error) {
	h := &AuthMdlwHdl{
		authService: authService,
		logger:      logger,
		reqCtxHdl:   reqCtxHdl,
	}
	if h.authService == nil {
		return nil, errors.New("nil auth service")
	}
	if h.logger == nil {
		return nil, errors.New("nil logger")
	}
	if h.reqCtxHdl == nil {
		return nil, errors.New("nil request context handler")
	}
	return h, nil
}

// Auth returns a middleware that checks the Authorization header for a valid token
// using the given authService. Depending on the authType, it can validate either an API key or a JWT token.
// If the token is valid, it sets the client ID and environment in the request context.
//
// Parameters:
//   - authType: The type of authentication to perform (API key or JWT token).
//
// Returns:
//   - A fiber.Handler that performs the authentication check and sets the request context.
func (h *AuthMdlwHdl) Auth(authType domain.AuthType) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("Auth middleware")
		// Get the Authorization header
		authHeader := c.Get("Authorization")

		// Check if it contains the Bearer prefix
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized, missing or invalid Authorization header",
			})
		}

		// Extract the token from the header, notice that this could be
		// an API key or a JWT token depending on the authType.
		token := strings.TrimPrefix(authHeader, "Bearer ")

		var authResult *domain.AuthResult
		var authErr error

		switch authType {
		case domain.AuthTypeAPIKey:
			authResult, authErr = h.authService.AuthWithAPIKey(c.Context(), token)
		case domain.AuthTypeToken:
			authResult, authErr = h.authService.AuthWithToken(c.Context(), token)
		default:
			h.logger.Error(c.Context(), "Invalid auth type", "authType", authType)
			authErr = errors.New("invalid auth type")
		}
		if authErr != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid API Key or token"})
		}

		h.reqCtxHdl.SetClientID(c, authResult.ClientID)
		h.reqCtxHdl.SetEnvironment(c, authResult.Environment)

		return c.Next()
	}
}
