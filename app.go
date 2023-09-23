package main

import "github.com/gofiber/fiber/v2"

func main() {
  app := fiber.New()
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("GO GO A FIBER!!!")
  })
  app.Listen(":3000")
}
