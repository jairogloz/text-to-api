package request_context_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
	"text-to-api/internal/domain"
	"text-to-api/internal/handlers/request_context"
)

func TestHandler_SetClientID_GetClientID(t *testing.T) {

	t.Run("non-empty clientID set", func(subTest *testing.T) {
		app := fiber.New()

		fasthttpCtx := &fasthttp.RequestCtx{}
		c := app.AcquireCtx(fasthttpCtx)

		h := request_context.NewRequestContextHandler()
		h.SetClientID(c, "client_id")

		clientId := h.GetClientID(c)
		assert.Equal(subTest, "client_id", clientId)

		app.ReleaseCtx(c)
	})

	t.Run("no clientID set", func(subTest *testing.T) {
		app := fiber.New()

		fasthttpCtx := &fasthttp.RequestCtx{}
		c := app.AcquireCtx(fasthttpCtx)

		h := request_context.NewRequestContextHandler()

		clientId := h.GetClientID(c)
		assert.Empty(subTest, clientId)

		app.ReleaseCtx(c)
	})

}

func TestHandler_SetEnvironment_GetEnvironment(t *testing.T) {

	t.Run("non-empty environment set", func(subTest *testing.T) {
		app := fiber.New()

		fasthttpCtx := &fasthttp.RequestCtx{}
		c := app.AcquireCtx(fasthttpCtx)

		h := request_context.NewRequestContextHandler()
		h.SetEnvironment(c, domain.RequestEnvironmentLive)

		requestEnvironment := h.GetEnvironment(c)
		assert.Equal(subTest, domain.RequestEnvironmentLive, requestEnvironment)

		app.ReleaseCtx(c)
	})

	t.Run("no environment set", func(subTest *testing.T) {
		app := fiber.New()

		fasthttpCtx := &fasthttp.RequestCtx{}
		c := app.AcquireCtx(fasthttpCtx)

		h := request_context.NewRequestContextHandler()

		env := h.GetEnvironment(c)
		assert.Empty(subTest, env)

		app.ReleaseCtx(c)
	})

}

func TestHandler_SetUserID_GetUserID(t *testing.T) {

	t.Run("non-empty userId set", func(subTest *testing.T) {
		app := fiber.New()

		fasthttpCtx := &fasthttp.RequestCtx{}
		c := app.AcquireCtx(fasthttpCtx)

		h := request_context.NewRequestContextHandler()
		h.SetUserID(c, "user_id")

		userId := h.GetUserID(c)
		assert.Equal(subTest, "user_id", userId)

		app.ReleaseCtx(c)
	})

	t.Run("no userId set", func(subTest *testing.T) {
		app := fiber.New()

		fasthttpCtx := &fasthttp.RequestCtx{}
		c := app.AcquireCtx(fasthttpCtx)

		h := request_context.NewRequestContextHandler()

		userID := h.GetUserID(c)
		assert.Empty(subTest, userID)

		app.ReleaseCtx(c)
	})

}
