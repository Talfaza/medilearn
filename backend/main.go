package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Welcome to Medilearn backend!")
    })

    app.Listen("0.0.0.0:3000")
}

