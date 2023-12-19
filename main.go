package main

import (
	"log"

	"musematch/controllers"
	"musematch/globals"
	"musematch/middleware"
	"musematch/queries"
	"musematch/routes"
	"musematch/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/template/html/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error
	err = globals.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

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

	err = utils.InitLog(globals.Env.LOG_FILE)
	if err != nil {
		log.Fatal("Failed to init log")
	}

	// init template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:        engine,
		ErrorHandler: controllers.ErrorController,
	})

	app.Use(middleware.Logger)
	app.Use("/", middleware.ThemeFromCookie)
	app.Get("/metrics", middleware.AdminProtected, monitor.New())

	routes.PrivateRoutes(app)
	routes.PublicRoutes(app)
	routes.AdminRoutes(app)

	app.Static("/", "./public")
	app.Use(controllers.NotFoundController)

	err = app.Listen(":3000")

	if err != nil {
		log.Fatal("Failed to listen port")
	}
}
