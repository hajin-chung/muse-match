package main

import (
	"log"

	"musematch/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// init db
	db, err := sqlx.Connect("sqlite3", "db/test.db")
	if err != nil {
		log.Fatal("Failed connecting db")
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use("/", logger.New())
	app.Use("/", func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	routes.PrivateRoutes(app)
	routes.PublicRoutes(app)

	app.Static("/", "./public")

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("Failed to listen port 3000")
	}
}
