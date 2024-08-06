package controllers

import "github.com/gofiber/fiber/v2"

func PostsIndex(c *fiber.Ctx) error {

	// what page we want the frontend to show
	return c.Render("posts/index", fiber.Map{})
}
