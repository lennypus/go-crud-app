package service

import (
  "github.com/husnulhamidiah/go-crud-app/contract"
  "github.com/husnulhamidiah/go-crud-app/model"
  "github.com/husnulhamidiah/go-crud-app/repository"
)

func GetAllBooks() []contract.Book {
  books := repository.GetAllBooks()

  var booksResponse []contract.Book
  for _, v := range books {
    b := contract.Book{
      Title: v.Title,
    }
    booksResponse = append(booksResponse, b)
  }
  return booksResponse
}

func GetBookById(id string) model.Book {
  // 3. panggil repository
  return repository.GetBookById(id)
}

