package main

import (
  "cideclasse/controllers"

  "github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New()

  controllers.DefineStudentsEndPoints(app)

  app.Listen(":3000")
}

