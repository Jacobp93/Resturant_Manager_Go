package main

import (
	"Resturant_Manager_Go/Resturant_Backend_Management/database"
	"Resturant_Manager_Go/Resturant_Backend_Management/middleware"
	"Resturant_Manager_Go/Resturant_Backend_Management/routes"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodcollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authenaction())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderitemRoutes(router)
	routes.InvoiceRotes(router)

	router.Run(":" + port)

}
