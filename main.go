package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello World",
			"Button": fiber.Map{
				"Href": "http://google.com",
				"Text": "hi",
			},
		})
	})

	app.Get("/art/:artId", func(c *fiber.Ctx) error {
		artId := c.Params("artId")
		return c.Render("index", fiber.Map{
			"Title": artId,
		})
	})

	app.Static("/", "./public")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Failed to listen port 3000")
	}
	log.Print("Listening to port 3000")
}
