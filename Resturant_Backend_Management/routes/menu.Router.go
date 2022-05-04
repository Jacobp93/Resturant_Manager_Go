package routes

import (
	controller "RESTURANT_MANAGER_GO/Resturant_Backend_Management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/Menus", controller.GetMenu())
	incomingRoutes.GET("/Menus/:Menu_id", controller.GetMenu())
	incomingRoutes.POST("/Menus", controller.CreateMenu())
	incomingRoutes.PATCH("/Menus/:Menu_id", controller.UpdateMenu())
}
