package main

import (
  "github.com/gofiber/fiber/v2"
  // "gorm.io/gorm"
  // "gorm.io/driver/mysql"
  // "fmt"
)

// type Login struct {
//   gorm.Model
//   Identificator  string
//   PasswordHash string
// }

func main() {
  // dsn := "root:mysql@tcp(127.0.0.1:3306)/mysqldb?charset=utf8mb4&parseTime=True&loc=Local"
  // db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  // if err != nil {
  //   return
  // }
  //
  // db.AutoMigrate(&Login{})
  //
  // db.Create(&Login{Identificator: "D42", PasswordHash: "dwa"})
  //
  // var login Login
  // db.First(&login, 1) // find product with integer primary key
  //
  // fmt.Println(login)

  app := fiber.New()

  app.Static("/", "./frontend/src")
  app.Static("dist/bundle.js", "./frontend/dist/bundle.js")
  app.Static("dist/bundle.css", "./frontend/dist/bundle.css")

  // app.Get("/", func(c *fiber.Ctx) error {
  //   return c.SendString("Hello, World!")
  // })

  app.Listen(":3000")
}

