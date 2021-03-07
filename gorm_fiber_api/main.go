package main

import (
  "fmt"

  "gorm_fiber_api/book"
  "gorm_fiber_api/database"

  "github.com/gofiber/fiber"
  "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(c *fiber.Ctx) {
  c.Send("Hello,  World")
}

func setupRoutes(app *fiber.App) {
  defaultRoute := "/api/v1/"

  app.Get(defaultRoute + "book", book.GetBooks)
  app.Get(defaultRoute + "book/:id", book.GetBook)
  app.Post(defaultRoute + "book", book.NewBook)
  app.Delete(defaultRoute + "book/:id", book.DeleteBook)
}

func initDatabase() {
  var err error

  database.DBconn, err = gorm.Open("sqlite3", "books.db")

  if err != nil {
    panic("Failed to connect to database")
  }

  fmt.Println("Database connection successfully opened")

  database.DBconn.AutoMigrate(&book.Book{})

  fmt.Println("Database migrated successfully")
}

func main() {
  app := fiber.New()

  initDatabase()
  defer database.DBconn.Close()

  app.Get("/", helloWorld)

  setupRoutes(app)

  app.Listen(3000)
}
