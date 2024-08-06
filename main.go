package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"go-mvc/initializers"
	"go-mvc/routes"
	"os"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	// Load Templates
	engine := html.New("./view", ".gohtml")

	// Setup views with the app
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	// app config
	//app.Static("/", "./public")

	// this is to pass the cookie to the frontend
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	// gets routes
	routes.Routes(app)

	app.Listen(":" + os.Getenv("PORT"))
}
