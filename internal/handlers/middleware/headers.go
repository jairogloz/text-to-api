package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"text-to-api/internal/domain"
	"text-to-api/internal/handlers"
	hdlPorts "text-to-api/internal/handlers/ports"
	"text-to-api/internal/ports"
)

// HeadersMdlwHdl is the receiver for the headers middleware. It holds the
// required components to perform the middleware.
type HeadersMdlwHdl struct {
	logger    ports.Logger
	reqCtxHdl hdlPorts.RequestContextHandler
}

// NewHeadersMdlwHdl returns a new instance of HeadersMdlwHdl. It returns an error if any of the
// required components are nil.
func NewHeadersMdlwHdl(l ports.Logger, reqCtxHdl hdlPorts.RequestContextHandler) (*HeadersMdlwHdl, error) {
	h := &HeadersMdlwHdl{
		logger:    l,
		reqCtxHdl: reqCtxHdl,
	}
	if h.logger == nil {
		return nil, fmt.Errorf("nil logger")
	}
	if h.reqCtxHdl == nil {
		return nil, fmt.Errorf("nil request context handler")
	}
	return h, nil
}

// ForceHeaders is a middleware that reads specified headers and stores them in the request context.
// It validates the presence and correctness of the "User-Id" and "Environment" headers.
//
// Parameters:
//   - headers: A slice of strings representing the headers to be processed.
//
// Returns:
//   - A fiber.Handler that processes the headers and sets the request context accordingly.
func (h *HeadersMdlwHdl) ForceHeaders(headers []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("ForceHeaders middleware")

		for _, header := range headers {
			switch header {
			case "User-Id":
				userID := c.Get("User-Id")
				if userID == "" {
					httpStatusCode, msg := handlers.ToHTTPError(fmt.Errorf("%w: missing User-Id header", domain.ErrorValidation))
					return c.Status(httpStatusCode).JSON(fiber.Map{
						"error": msg,
					})
				}
				h.reqCtxHdl.SetUserID(c, userID)
			case "Environment":
				e := c.Get("Environment")
				environment, err := domain.NewRequestEnvironment(e)
				if err != nil {
					httpStatusCode, msg := handlers.ToHTTPError(err)
					return c.Status(httpStatusCode).JSON(fiber.Map{
						"error": msg,
					})
				}
				h.reqCtxHdl.SetEnvironment(c, environment)
			default:
				h.logger.Error(c.Context(), "Unknown header: %s", header)
				httpStatusCode, msg := handlers.ToHTTPError(errors.New("unknown header"))
				return c.Status(httpStatusCode).JSON(fiber.Map{
					"error": msg,
				})
			}
		}

		return c.Next()
	}
}
