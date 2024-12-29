package main

import (
	"log"
	"os"

	"github.com/alperencantez/ignite/database"
	"github.com/joho/godotenv"

	"github.com/alperencantez/ignite/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	var envFile string

	switch env {
	case "development":
		envFile = ".env.development"
	case "production":
		envFile = ".env.production"
	default:
		log.Fatalf("Unknown environment: %s", env)
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file: %v", envFile, err)
	}
}

func main() {
	database.ConnectToDb()
	app := fiber.New(fiber.Config{
		AppName: "Ignite v1.0",
	})

	app.Use(cors.New())
	router.InitRoutes(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
