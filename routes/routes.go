package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-mvc/controllers"
)

// pass in app to the setup routes function
func Routes(app *fiber.App) {
	//GET  -------------------------------------------------
	app.Get("/api/user", controllers.User)

	//POST -------------------------------------------------
	app.Post("/api/register", controllers.Register) // create the path in the controllers folder
	app.Post("/api/Login", controllers.Login)       // API path to create the Auth

	// PUT  -------------------------------------------------

	// DELETE -----------------------------------------------
}
