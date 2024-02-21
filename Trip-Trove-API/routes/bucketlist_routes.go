package routes

import (
	"Trip-Trove-API/presentation/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterBucketListRoutes(router *gin.Engine, bucketListHandler *handlers.BucketListHandler) {
	locationGroup := router.Group("/bucketList")
	{
		locationGroup.GET("/", bucketListHandler.AllBucketLists)
		locationGroup.GET("/:id", bucketListHandler.BucketListByID)
		locationGroup.GET("/user/:userId", bucketListHandler.BucketListsByUserID)
		locationGroup.POST("/", bucketListHandler.CreateBucketList)
		locationGroup.PUT("/:id", bucketListHandler.UpdateBucketList)
		locationGroup.DELETE("/:id", bucketListHandler.DeleteBucketList)
	}
}
