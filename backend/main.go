package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/little-fox28/React-Go/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func main() {

	// DB connection
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .evn file: ", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB...")
	collection = client.Database("react-go").Collection("todos")

	// Server
	app := fiber.New()
	app.Use(cors.New())
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	routes.Route(app, collection)

	log.Fatal(app.Listen(":" + PORT))

}
