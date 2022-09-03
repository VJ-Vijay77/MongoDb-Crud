package main

import (
	"os"

	"github.com/VJ-Vijay77/mongodb/config"
	"github.com/VJ-Vijay77/mongodb/db"
	"github.com/VJ-Vijay77/mongodb/errpkg"
	"github.com/VJ-Vijay77/mongodb/routes"

	"github.com/gofiber/fiber/v2"
)



func main() {


	//! setting the fiber Engine
	api := fiber.New()



	//? loading env files
	config.LoadEnv()
	
	//! getting necessary variable data from .env
	MongoURI := os.Getenv("MONGO_URI")
	Port := os.Getenv("PORT")



	//? connecting to mongodb instance
	client, err := db.Connect(MongoURI)
	errpkg.Fatal(err)
	//closing the channel when all resource ended
	// defer cancel()
	_ = client


	//? calling the routes
	routes.FiberRoutes(api)

	//connecting to port
	api.Listen(":"+Port)
	

}
