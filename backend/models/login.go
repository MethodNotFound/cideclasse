package models

import (
  "gorm.io/gorm"
)

type Login struct {
 gorm.Model
  Identificator string `gorm:"not null;index:,unique"`
  PasswordHash string `gorm:"not null"`

  StudentID int `gorm:"not null"`
  Student Student
}
