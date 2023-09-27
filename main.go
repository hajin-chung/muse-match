package main

import (
	"log"

	"musematch/app/globals"
	"musematch/app/middleware"
	"musematch/app/queries"
	"musematch/app/routes"
	"musematch/app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error
	err = globals.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", globals.Env)

	err = queries.InitDB()
	if err != nil {
		log.Fatal("Failed to init db")
	}

	err = utils.InitS3()
	if err != nil {
		log.Fatal("Failed to init s3")
	}

	// create session store
	globals.InitStore()

	// init template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use("/", logger.New())
	app.Use("/", middleware.ThemeFromCookie)

	routes.PrivateRoutes(app)
	routes.PublicRoutes(app)
	routes.AdminRoutes(app)

	app.Static("/", "./public")

	err = app.Listen(":3000")

	if err != nil {
		log.Fatal("Failed to listen port")
	}
}
