package repository

import (
  "github.com/husnulhamidiah/go-crud-app/database"
  "github.com/husnulhamidiah/go-crud-app/model"
)

func GetAllBooks() []model.Book {
  var books []model.Book
  database.DB.Find(&books)
  return books
}

func GetBookById(id string) model.Book {
  // 4. call database
  var book model.Book
  database.DB.First(&book, id)
  return book
}