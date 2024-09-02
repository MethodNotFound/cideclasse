package models

import (
  "gorm.io/gorm"
)

type Class struct {
  gorm.Model
  Name  string `gorm:"not null;index:,unique"`

  Students []*Student `gorm:"many2many:user_classes;"`
}

func CreateClass(db *gorm.DB, name string) (*Class, error) {
  class := Class{Name: name}

  err := db.Create(&class).Error
  if err != nil {
    return nil, err
  }
  return &class, nil

}

func FindClass(db *gorm.DB, name string) (*Class, error){
    var class Class

    result := db.Where("name = ?", name).First(&class)
    if result.Error != nil {
        return nil, result.Error
    }
    return &class, nil
}
