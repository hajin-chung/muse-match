package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db/test.db")
	if err != nil {
		log.Fatal("database missing")
	}
	defer db.Close()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT id, name FROM user")
		if err != nil {
			log.Println("error on selecting from user")
		}
		defer rows.Close()

		for rows.Next() {
			var id string
			var name string

			err = rows.Scan(&id, &name)
			if err != nil {
				log.Println("error on row scan", err)
			}
			log.Println(id, name)
		}
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
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("Failed to listen port 3000")
	}
	log.Print("Listening to port 3000")
}
