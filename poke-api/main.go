package main

import (
  "poke-api/poke"
  "github.com/gofiber/fiber"
)

func pong(c *fiber.Ctx) {
  c.Send("pong")
}

func setupRoutes(app *fiber.App) {
  defaultRoute := "/api/v1/"

  app.Get(defaultRoute + "pokemon/:name", poke.GetPokemon)
}

func main () {
  app := fiber.New()

  app.Get("/ping", pong)

  setupRoutes(app)

  app.Listen(3000)
}
