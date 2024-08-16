package main

import "github.com/gofiber/fiber/v2"

func main() {
  app := fiber.New()

  app.Static("/", "./frontend/src")
  app.Static("/Bundle.js", "./frontend/Bundle.js")

  // app.Get("/", func(c *fiber.Ctx) error {
  //   return c.SendString("Hello, World!")
  // })

  app.Listen(":3000")
}





