package routes

import (
	"github.com/RafaelMoreira96/secure-campus-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Rotas públicas

	// Rotas protegidas
	app.Use(utils.JWTMiddleware)
}
