package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-restful-api/controllers"
)

func AuthorRoutes(api fiber.Router) {
	router := api.Group("/authors")
	router.Post("/", controllers.CreateAuthor)
	router.Get("/", controllers.GetAuthors)
	router.Get("/:id", controllers.GetAuthorByID)
	router.Patch("/:id", controllers.UpdateAuthor)
	router.Delete("/:id", controllers.DeleteAuthor)

}
