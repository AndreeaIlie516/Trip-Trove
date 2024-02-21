package handlers

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/domain/services"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type InBucketListHandler struct {
	Service *services.InBucketListService
}

func (handler *InBucketListHandler) AllInBucketListAssociations(c *gin.Context) {
	inBucketListAssociations, err := handler.Service.AllInBucketListAssociations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch in bucket list associations"})
		return
	}
	c.JSON(http.StatusOK, inBucketListAssociations)
}

func (handler *InBucketListHandler) InBucketListAssociationByID(c *gin.Context) {
	id := c.Param("id")
	inBucketListAssociation, err := handler.Service.InBucketListAssociationByID(id)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "in bucket list association not found"})
		}
		return
	}
	c.JSON(http.StatusOK, inBucketListAssociation)
}

func (handler *InBucketListHandler) InBucketListAssociation(c *gin.Context) {
	destinationID := c.Query("destinationId")
	bucketListID := c.Query("bucketListId")

	if destinationID == "" || bucketListID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "destinationId and bucketListId query parameters are required"})
		return
	}

	inBucketListAssociation, err := handler.Service.InBucketListAssociation(destinationID, bucketListID)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "in bucket list association not found"})
		}
		return
	}
	c.JSON(http.StatusOK, inBucketListAssociation)
}

func (handler *InBucketListHandler) DestinationWithBucketLists(c *gin.Context) {
	destinationID := c.Param("destinationId")
	destinationWithBucketLists, err := handler.Service.DestinationWithBucketLists(destinationID)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, destinationWithBucketLists)
}

func (handler *InBucketListHandler) BucketListWithDestinations(c *gin.Context) {
	bucketListID := c.Param("bucketListid")

	bucketListWithDestinations, err := handler.Service.BucketListWithDestinations(bucketListID)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, bucketListWithDestinations)
}

func (handler *InBucketListHandler) AddDestinationToBucketList(c *gin.Context) {
	var newInBucketListAssociation entities.InBucketList

	if err := c.BindJSON(&newInBucketListAssociation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.Struct(newInBucketListAssociation)

	if err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid validation error"})
			return
		}

		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := "Validation error on field '" + err.Field() + "': " + err.ActualTag()
			if err.Param() != "" {
				errorMessage += " (Parameter: " + err.Param() + ")"
			}
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	inBucketListAssociation, err := handler.Service.AddDestinationToBucketList(newInBucketListAssociation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, inBucketListAssociation)
}

func (handler *InBucketListHandler) DeleteDestinationFromBucketList(c *gin.Context) {
	id := c.Param("id")

	inBucketListAssociation, err := handler.Service.DeleteDestinationFromBucketList(id)

	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, inBucketListAssociation)
}
