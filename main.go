package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-restful-api/config"
	"go-restful-api/routes"
	"log"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()
	log.Println("Server running on port 8080")

	app := fiber.New()
	routes.SetupRoutes(app)
	if err := app.Listen(fmt.Sprintf(":%v", config.ENV.PORT)); err != nil {
		log.Fatal(err)
	}

}
