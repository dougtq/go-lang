package book

import (
  "github.com/gofiber/fiber"
  "github.com/jinzhu/gorm"
  "gorm_fiber_api/database"
)

type Book struct {
  gorm.Model
  Title string `json:"title"`
  Author string `json:"author"`
  Rating int `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
  db := database.DBconn

  var books []Book

  db.Find(&books)

  c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
  db := database.DBconn

  id := c.Params("id")

  var book Book

  db.Find(&book, id)

  c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
  db := database.DBconn

  book := new(Book)

  err := c.BodyParser(book)

  if err != nil {
    c.Status(503).Send(err)
    return
  }

  db.Create(&book)
  c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
  db := database.DBconn

  id := c.Params("id")

  var book Book

  db.First(&book, id)

  if book.Title == "" {
    c.Status(500).Send("No book found with the given id")
  }

  db.Delete(&book)

  c.Send("Book successfully deleted")
}
