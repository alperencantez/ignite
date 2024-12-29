package router

import (
	"github.com/alperencantez/ignite/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())

	// Version 1.x
	v1 := api.Group("/v1")
	v1.Get("/test", handler.Hello)
}
