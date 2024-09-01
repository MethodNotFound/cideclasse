package models

import (
  "gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Session struct {
 gorm.Model
  Active    bool `gorm:"not null"`
  MetaData  map[string]interface{} `gorm:"serializer:json"`
  StudentID int `gorm:"not null"`
  Student   Student
}

func NewSession(db *gorm.DB, identifier string, password string) (*Session, error){
  student, err := FindStudent(db, identifier)
  if err != nil {
    return nil, err
  }

  if err := bcrypt.CompareHashAndPassword([]byte(student.PasswordHash), []byte(password)); err != nil {
    return nil, err
  } else {
    session := Session{Active: true, MetaData: nil, Student: *student}
    err := db.Create(&session).Error
    if err != nil {
      return nil, err
    }
    return &session, nil
  }
}
