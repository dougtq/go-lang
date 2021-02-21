package main

import (
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

type Todo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{
	{ID: 1, Name: "Learn golang", Completed: false},
	{ID: 2, Name: "Learn rust", Completed: false},
	{ID: 3, Name: "Learn elixir", Completed: false},
}

func main() {
	app := fiber.New()

	app.Use(middleware.Logger())

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("hello")
	})

	todosRoute := app.Group("/todos")
	todosRoute.Get("/", func(ctx *fiber.Ctx) {
		ctx.Status(fiber.StatusOK).JSON(todos)
	})
	todosRoute.Post("/", createTodo)
	todosRoute.Get("/:id", getTodo)
	todosRoute.Delete("/:id", deleteTodo)
	todosRoute.Patch("/:id", updateTodo)

	err := app.Listen(3000)

	if err != nil {
		panic(err)
	}

}

func createTodo(ctx *fiber.Ctx) {
	type request struct {
		Name string `json:"name"`
	}

	var body request
	err := ctx.BodyParser(&body)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return
	}

	todo := &Todo{
		ID:        len(todos) + 1,
		Name:      body.Name,
		Completed: false,
	}

	todos = append(todos, todo)

	ctx.Status(fiber.StatusCreated).JSON(todo)
}

func getTodo(ctx *fiber.Ctx) {
	paramsID := ctx.Params("id")
	id, err := strconv.Atoi(paramsID)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id could not be parsed",
		})
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			ctx.Status(fiber.StatusOK).JSON(todo)
			return
		}
	}

	ctx.Status(fiber.StatusNotFound)
	return
}

func deleteTodo(ctx *fiber.Ctx) {
	paramsID := ctx.Params("id")
	id, err := strconv.Atoi(paramsID)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id could not be parsed",
		})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[0:i], todos[i+1:]...)
			ctx.Status(fiber.StatusNoContent).JSON(todos)
			return
		}
	}

	ctx.Status(fiber.StatusNotFound)
	return
}

func updateTodo(ctx *fiber.Ctx) {
	type request struct {
		Name      *string `json:"name"`
		Completed *bool   `json:"completed"`
	}

	paramsID := ctx.Params("id")
	id, err := strconv.Atoi(paramsID)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id could not be parsed",
		})
		return
	}

	var body request

	err = ctx.BodyParser(&body)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse body",
		})
		return
	}

	var todo *Todo

	for _, t := range todos {
		if t.ID == id {
			todo = t
			break
		}
	}

	if todo == nil {
		ctx.Status(fiber.StatusNotFound)
		return
	}

	if body.Name != nil {
		todo.Name = *body.Name
	}

	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	ctx.Status(fiber.StatusOK).JSON(todo)

	return
}
