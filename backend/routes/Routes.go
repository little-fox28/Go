package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/little-fox28/React-Go/models"
)

var todos = []models.Todo{};

func getTodo (c *fiber.Ctx) error {
	return c.Status(200).JSON(todos)
}

func addTodo (c *fiber.Ctx) error {
	todo := &models.Todo{}

	// Validate the body
	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" || len(todo.Body) < 2 {
		return c.Status(400).JSON(fiber.Map{"msg": "Todo body is required"})
	}

	todo.ID = len(todos) + 1
	todos = append(todos, *todo)

	return c.Status(201).JSON(todos)
}

func updateTodo (c *fiber.Ctx) error {
	id := c.Params("id")

	for i, todo := range todos {
		if fmt.Sprint(todo.ID) == id {
            todos[i].Completed = !todo.Completed			
			return c.Status(200).JSON(todos[i])
        }
	}

	return c.Status(400).JSON(fiber.Map{"msg": "Todo not found"})
}

func deleteTodo (c *fiber.Ctx) error {
	id := c.Params("id")
	
	for i,todo := range todos {
		if fmt.Sprint(todo.ID) == id {
		    todos = append(todos[:i],todos[i+1:]...)

		    return c.Status(200).JSON(fiber.Map{"success": true})
		}
	}

    return c.Status(400).JSON(fiber.Map{"msg":"Todo not found"})
}


func Route(app *fiber.App) {
	api := app.Group("/apis/v1")

	api.Get("/todo", getTodo)

	api.Post("/todo", addTodo)

	api.Patch("/todo/:id", updateTodo)

	api.Delete("/todo/:id", deleteTodo)

}
