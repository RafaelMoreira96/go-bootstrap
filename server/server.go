package server

import (
	"github.com/RafaelMoreira96/secure-campus-backend/database"
	"github.com/RafaelMoreira96/secure-campus-backend/server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func RunServer() {
	app := fiber.New()

	// Segurança
	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8000, https://localhost:8000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Use(compress.New())

	// Banco de dados e rotas
	database.Connect()
	routes.SetupRoutes(app)

	app.Listen(":8000")
}
