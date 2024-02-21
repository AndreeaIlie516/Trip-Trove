package routes

import (
	"Trip-Trove-API/presentation/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterCityRoutes(router *gin.Engine, cityHandler *handlers.CityHandler) {
	cityGroup := router.Group("/cities")
	{
		cityGroup.GET("/", cityHandler.AllCities)
		cityGroup.GET("/:id", cityHandler.CityByID)
		cityGroup.POST("/", cityHandler.CreateCity)
		cityGroup.PUT("/:id", cityHandler.UpdateCity)
		cityGroup.DELETE("/:id", cityHandler.DeleteCity)
	}
}
