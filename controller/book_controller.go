package controller

import (
  "github.com/gofiber/fiber/v2"
  "github.com/husnulhamidiah/go-crud-app/contract"
  "github.com/husnulhamidiah/go-crud-app/service"
)

func GetAllBooks (c *fiber.Ctx) error {
  var books []contract.Book
  books = service.GetAllBooks()
  return c.JSON(books)
}

func GetBookById (c *fiber.Ctx) error {
  // 2. panggil service
  id := c.Params("id")
  book := service.GetBookById(id)
  return c.JSON(book)
}