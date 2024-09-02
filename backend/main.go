package main

import (
  "cideclasse/controllers"

  "github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
	// "github.com/golang-jwt/jwt/v5"
)

func main() {
  app := fiber.New()

  controllers.DefineSessionsEndPoints(app)

  app.Use(jwtware.New(jwtware.Config{
    SigningKey: jwtware.SigningKey{Key: []byte("secret")},
  }))

  controllers.DefineAdminStudentsEndPoints(app)
  controllers.DefineAdminClassesEndPoints(app)
  controllers.DefineStudentsEndPoints(app)

  app.Listen(":3000")
}

