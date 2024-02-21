package routes

import (
	"Trip-Trove-API/presentation/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterInBucketListRoutes(router *gin.Engine, inBucketListHandler *handlers.InBucketListHandler) {
	eventArtistGroup := router.Group("/favourite-events")
	{
		eventArtistGroup.GET("/", inBucketListHandler.AllInBucketListAssociations)
		eventArtistGroup.GET("/:id", inBucketListHandler.InBucketListAssociationByID)
		eventArtistGroup.GET("/associationByEventAndUser", inBucketListHandler.InBucketListAssociation)
		eventArtistGroup.GET("/destination/:destinationId", inBucketListHandler.DestinationWithBucketLists)
		eventArtistGroup.GET("/bucketList/:bucketListId", inBucketListHandler.BucketListWithDestinations)
		eventArtistGroup.POST("/", inBucketListHandler.AddDestinationToBucketList)
		eventArtistGroup.DELETE("/:id", inBucketListHandler.DeleteDestinationFromBucketList)
	}
}
