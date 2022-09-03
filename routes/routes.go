package routes

import (
	"github.com/VJ-Vijay77/mongodb/controllers"
	"github.com/gofiber/fiber/v2"
)



func FiberRoutes(api *fiber.App) error {
	api.Get("/person/:id?",controllers.GetPerson)
	api.Post("/person",controllers.CreatePerson)
	api.Put("/person/:id",controllers.UpdatePerson)
	api.Delete("/person/:id",controllers.DeletePerson)
	
	return nil
}