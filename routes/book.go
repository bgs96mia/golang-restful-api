package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-restful-api/controllers"
)

func BookRoutes(api fiber.Router) {
	router := api.Group("/books")
	router.Post("/", controllers.CreateBook)
	router.Get("/", controllers.GetBooks)
	router.Get("/:id", controllers.GetBookByID)
	router.Patch("/:id", controllers.UpdateBook)
	router.Delete("/:id", controllers.DeleteBook)

}
