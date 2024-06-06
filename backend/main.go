package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/little-fox28/React-Go/routes"
)


func main() {
    app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .evn file")
	}

	PORT := os.Getenv("PORT")

	routes.Route(app)

    log.Fatal(app.Listen(":" + PORT))
}