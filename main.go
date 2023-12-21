package main

import (
	"log"

	"musematch/controllers"
	"musematch/globals"
	"musematch/middleware"
	"musematch/queries"
	"musematch/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
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

	app := fiber.New(fiber.Config{
		// ErrorHandler: controllers.ErrorController,
		// TODO: create error handler
	})

	app.Use(middleware.Logger)
	app.Get("/metrics", middleware.AdminProtected, monitor.New())
	app.Static("/", "./public")

	app.Use(middleware.ContentTypeHtml)
	app.Get("/", controllers.MainController)
	app.Get("/auth/login", controllers.LoginController)
	app.Get("/auth/callback/kakao", controllers.KakaoCallbackController)

	app.Get("/dashboard/arts", middleware.SessionProtected, controllers.DashboardArtController)
	// app.Get("/auth/callback/naver", controllers.NaverCallbackController)

	// app.Use(controllers.NotFoundController)

	err = app.Listen(":3000")

	if err != nil {
		log.Fatal("Failed to listen port")
	}
}
