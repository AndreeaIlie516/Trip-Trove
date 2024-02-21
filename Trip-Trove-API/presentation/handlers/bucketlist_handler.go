package handlers

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/domain/services"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type BucketListHandler struct {
	Service *services.BucketListService
}

func (handler *BucketListHandler) AllBucketLists(c *gin.Context) {
	bucketLists, err := handler.Service.AllBucketLists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch bucketLists"})
		return
	}
	c.JSON(http.StatusOK, bucketLists)
}

func (handler *BucketListHandler) BucketListByID(c *gin.Context) {
	id := c.Param("id")
	bucketList, err := handler.Service.BucketListByID(id)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, bucketList)
}

func (handler *BucketListHandler) BucketListsByUserID(c *gin.Context) {
	userID := c.Param("userId")
	bucketLists, err := handler.Service.BucketListsByUserID(userID)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, bucketLists)
}

func (handler *BucketListHandler) CreateBucketList(c *gin.Context) {
	var newBucketList entities.BucketList

	if err := c.BindJSON(&newBucketList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.Struct(newBucketList)

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

	bucketList, err := handler.Service.CreateBucketList(newBucketList)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create bucketList"})
		return
	}

	c.JSON(http.StatusCreated, bucketList)
}

func (handler *BucketListHandler) DeleteBucketList(c *gin.Context) {
	id := c.Param("id")

	bucketList, err := handler.Service.DeleteBucketList(id)

	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, bucketList)
}

func (handler *BucketListHandler) UpdateBucketList(c *gin.Context) {
	id := c.Param("id")

	var updatedBucketList entities.BucketList

	if err := c.BindJSON(&updatedBucketList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.Struct(updatedBucketList)

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

	bucketList, err := handler.Service.UpdateBucketList(id, updatedBucketList)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bucketList not found"})
		return
	}

	c.JSON(http.StatusOK, bucketList)
}
