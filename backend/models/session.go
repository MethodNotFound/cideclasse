package models

import (
  "gorm.io/gorm"
)

type Session struct {
 gorm.Model
  Active    bool `gorm:"not null"`
  MetaData  map[string]interface{} `gorm:"serializer:json"`
  StudentID int `gorm:"not null"`
  Student   Student
}
