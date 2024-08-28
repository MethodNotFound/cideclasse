package models

import (
  "gorm.io/gorm"
)

type Student struct {
  gorm.Model
  Name  string `gorm:"not null"`
  Email string

  //login credentials
  Identifier string `gorm:"not null;index:,unique"` // rgm
  PasswordHash string
  AskNewPassword bool `gorm:"not null"`

  Sessions []Session
}

func CreateStudent(db *gorm.DB, name string, identifier string) (*Student, error){
  student := Student{Name: name, Identifier: identifier}

  err := db.Create(&student).Error
  if err != nil {
    return nil, err
  }
  return &student, nil
}
