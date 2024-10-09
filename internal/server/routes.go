package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) HealthHandler(c *fiber.Ctx) error {

	// return c.JSON(s.db.Health())
	// todo: Implement actual health checker
	return c.JSON(fiber.Map{"status": "ok"})

}
