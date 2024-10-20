package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {

    })

    app.Listen("0.0.0.0:3000")
}

