package request_context_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
	"text-to-api/internal/domain"
	"text-to-api/internal/handlers/request_context"
)

func TestHandler_SetClientID(t *testing.T) {
	app := fiber.New()

	t.Run("request context not initialized", func(subTest *testing.T) {
		// Create a new fasthttp.RequestCtx for each test
		fasthttpCtx := &fasthttp.RequestCtx{}
		c := app.AcquireCtx(fasthttpCtx)

		h := request_context.NewRequestContextHandler()
		// Run the function you want to test
		h.SetClientID(c, "client_id")

		requestContext, ok := c.Locals(domain.CtxKeyRequestContext).(*domain.RequestContext)
		assert.True(subTest, ok)
		if assert.NotEmpty(subTest, requestContext) {
			assert.Equal(subTest, "client_id", requestContext.ClientID)
		}

		// Release the context to avoid memory leaks
		app.ReleaseCtx(c)
	})

}
