package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"text-to-api/internal/handlers"
	hdlPorts "text-to-api/internal/handlers/ports"
	"text-to-api/internal/ports"
	"time"
)

// CheckSubscriptionMdlw is a middleware that checks if the client has an active subscription.
type CheckSubscriptionMdlw struct {
	logger    ports.Logger
	reqCtxHdl hdlPorts.RequestContextHandler
	subsSrv   ports.SubscriptionService
}

// NewCheckSubscriptionMdlw creates a new CheckSubscriptionMdlw. It returns an error if any of the required
// parameters is nil.
func NewCheckSubscriptionMdlw(l ports.Logger, r hdlPorts.RequestContextHandler,
	subsSrv ports.SubscriptionService) (*CheckSubscriptionMdlw, error) {
	c := &CheckSubscriptionMdlw{
		logger:    l,
		reqCtxHdl: r,
		subsSrv:   subsSrv,
	}
	if c.logger == nil {
		return nil, errors.New("nil logger")
	}
	if c.reqCtxHdl == nil {
		return nil, errors.New("nil request context handler")
	}
	if c.subsSrv == nil {
		return nil, errors.New("nil subscription service")
	}
	return c, nil
}

// CheckSubscription is a middleware that verifies the subscription status of the client making the request.
// It retrieves the client ID and environment from the request context, checks the subscription status,
// and sets the subscription information in the request context for downstream handlers.
//
// Returns:
//   - A Fiber handler function.
//
// Behavior:
//   - Extracts the `clientID` and `environment` from the request context using the `reqCtxHdl`.
//   - If `clientID` or `environment` is missing, logs an error and responds with a `500 Internal Server Error`.
//   - Calls the subscription service (`subsSrv`) to check the subscription status for the client.
//   - If the subscription check fails, converts the error into an appropriate HTTP status code and message, and responds with it.
//   - If successful, sets the subscription information in the request context using `reqCtxHdl`.
//   - Proceeds to the next middleware or handler in the chain if the subscription is valid.
func (ch *CheckSubscriptionMdlw) CheckSubscription() fiber.Handler {
	return func(c *fiber.Ctx) error {

		startCheckSubscription := time.Now().UTC()
		clientID := ch.reqCtxHdl.GetClientID(c)
		if clientID == "" {
			ch.logger.Error(c.Context(), "Client ID not found in request context")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
		environment := ch.reqCtxHdl.GetEnvironment(c)
		if environment == "" {
			ch.logger.Error(c.Context(), "Environment not found in request context")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}

		subscription, err := ch.subsSrv.CheckSubscription(c.Context(), clientID, environment)
		if err != nil {
			httpStatusCode, message := handlers.ToHTTPError(err)
			return c.Status(httpStatusCode).JSON(fiber.Map{"error": message})
		}

		ch.reqCtxHdl.SetSubscription(c, subscription)

		ch.logger.Debug(c.Context(), "Check subscription time", "time", time.Since(startCheckSubscription).String())

		return c.Next()
	}
}
