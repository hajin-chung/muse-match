package main

import (
	"log"

	"musematch/app/globals"
	"musematch/app/queries"
	"musematch/app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error
	// load env
	err = globals.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", globals.Env)

	// init db
	err = queries.InitDB()
	if err != nil {
		log.Fatal("Failed to init db")
	}

	// create session store
	globals.InitStore()

	// init template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use("/", logger.New())

	routes.PrivateRoutes(app)
	routes.PublicRoutes(app)

	app.Static("/", "./public")

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("Failed to listen port 3000")
	}
}
