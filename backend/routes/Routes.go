package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/little-fox28/React-Go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Route(app *fiber.App, collection *mongo.Collection) {

	api := app.Group("/apis/v1")

	api.Get("/todo", func(c *fiber.Ctx) error {
		return getTodo(c, collection)
	})

	api.Post("/todo/:id", func(c *fiber.Ctx) error {
		return getTodoByID(c, collection)
	})

	api.Post("/todo", func(c *fiber.Ctx) error {
		return createTodo(c, collection)
	})

	api.Patch("/todo/:id", func(c *fiber.Ctx) error {
		return updateTodo(c, collection)
	})

	api.Delete("/todo/:id", func(c *fiber.Ctx) error {
		return deleteTodo(c, collection)
	})

}

func getTodo(c *fiber.Ctx, collection *mongo.Collection) error {
	var todos []models.Todo

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}

	return c.Status(200).JSON(todos)
}

func getTodoByID(c *fiber.Ctx, collection *mongo.Collection) error {
	var todo models.Todo
	id := c.Params("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": objectId}

	err = collection.FindOne(context.Background(), filter).Decode(&todo)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	}

	return c.Status(200).JSON(todo)
}

func createTodo(c *fiber.Ctx, collection *mongo.Collection) error {
	todo := new(models.Todo)

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(fiber.Map{"success": true, "todo": todo})
}

func updateTodo(c *fiber.Ctx, collection *mongo.Collection) error {
	id := c.Params("id")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}

func deleteTodo(c *fiber.Ctx, collection *mongo.Collection) error {
	id := c.Params("id")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": objectId}

	_, err = collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})

}
