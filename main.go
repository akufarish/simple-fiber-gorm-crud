package main

import (
	muridController "fiber/controllers"
	"fiber/models"

	"github.com/gofiber/fiber/v2"
)

func main()  {
	Router := fiber.New()
	models.ConnectDatabase()

	Router.Get("/api/murid", muridController.Index)
	Router.Post("/api/murid", muridController.Store)
	Router.Get("/api/murid/:id", muridController.Show)
	Router.Put("/api/murid/:id", muridController.Update)
	Router.Delete("/api/murid/:id", muridController.Delete)

	Router.Listen(":8000")
}