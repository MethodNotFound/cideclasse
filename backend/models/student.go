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

  Classes []*Class `gorm:"many2many:user_classes;"`
}

func CreateStudent(db *gorm.DB, name string, identifier string) (*Student, error){
  student := Student{Name: name, Identifier: identifier, AskNewPassword: true}

  err := db.Create(&student).Error
  if err != nil {
    return nil, err
  }
  return &student, nil
}

func FindStudent(db *gorm.DB, identifier string) (*Student, error){
    var student Student

    result := db.Where("identifier = ?", identifier).First(&student)
    if result.Error != nil {
        return nil, result.Error
    }
    return &student, nil
}
