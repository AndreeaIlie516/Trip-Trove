package routes

import (
	"Trip-Trove-API/presentation/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterDestinationRoutes(router *gin.Engine, destinationHandler *handlers.DestinationHandler) {
	eventGroup := router.Group("/destinations")
	{
		eventGroup.GET("/", destinationHandler.AllDestinations)
		eventGroup.GET("/:id", destinationHandler.DestinationByID)
		eventGroup.GET("/location/:locationId", destinationHandler.DestinationsByLocationID)
		eventGroup.POST("/", destinationHandler.CreateDestination)
		eventGroup.PUT("/:id", destinationHandler.UpdateDestination)
		eventGroup.DELETE("/:id", destinationHandler.DeleteDestination)
	}
}
