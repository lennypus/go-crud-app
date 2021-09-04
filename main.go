package main

import (
  "fmt"
  "github.com/gofiber/fiber/v2"
  "github.com/husnulhamidiah/go-crud-app/controller"
  "github.com/husnulhamidiah/go-crud-app/database"
  "github.com/husnulhamidiah/go-crud-app/model"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "log"
)

func setupRoutes(app *fiber.App) {
  app.Get("/book", controller.GetAllBooks)

  // ROUTER -> CONTROLLER -> SERVICE -> REPOSITORY -> DB
  // 1. define router
  app.Get("/book/:id", controller.GetBookById)
}

func main() {
  app := fiber.New()

  var err error
  database.DB, err = gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/rakamin?parseTime=true"), &gorm.Config{})
  if err != nil {
    panic(err)
  }

  // migrasi table, hapus kalau tidak perlu
  database.DB.AutoMigrate(&model.Book{})
  fmt.Println("migrasi sukses")

  setupRoutes(app)
  log.Fatal(app.Listen(":3000"))
}

























