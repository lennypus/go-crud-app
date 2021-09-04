package model

import "gorm.io/gorm"

type Book struct {
  gorm.Model
  Author string
  Title string
  Rating uint
}
