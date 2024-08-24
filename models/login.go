package models

import (
  "gorm.io/gorm"
)

type Login struct {
  gorm.Model
  Identificator string `gorm:"not null"`
  PasswordHash string `gorm:"not null"`
  Type string `gorm:"not null"`
}
