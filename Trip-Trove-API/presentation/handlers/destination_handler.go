package handlers

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/domain/services"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type DestinationHandler struct {
	Service *services.DestinationService
}

func (handler *DestinationHandler) AllDestinations(c *gin.Context) {
	destinations, err := handler.Service.AllDestinations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch destinations"})
		return
	}
	c.JSON(http.StatusOK, destinations)
}

func (handler *DestinationHandler) DestinationByID(c *gin.Context) {
	id := c.Param("id")
	destination, err := handler.Service.DestinationByID(id)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "destination not found"})
		}
		return
	}
	c.JSON(http.StatusOK, destination)
}

func (handler *DestinationHandler) DestinationsByLocationID(c *gin.Context) {
	locationID := c.Param("locationId")
	destinations, err := handler.Service.DestinationsByLocationID(locationID)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, destinations)
}

func (handler *DestinationHandler) DestinationsByCityID(c *gin.Context) {
	cityID := c.Param("cityId")
	destinations, err := handler.Service.DestinationsByCityID(cityID)
	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, destinations)
}

func (handler *DestinationHandler) CreateDestination(c *gin.Context) {
	var newDestination entities.Destination

	if err := c.BindJSON(&newDestination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.Struct(newDestination)

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

	destination, err := handler.Service.CreateDestination(newDestination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create destination"})
		return
	}

	c.JSON(http.StatusCreated, destination)
}

func (handler *DestinationHandler) DeleteDestination(c *gin.Context) {
	id := c.Param("id")

	destination, err := handler.Service.DeleteDestination(id)

	if err != nil {
		if err.Error() == "invalid ID format" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "destination not found"})
		}
		return
	}

	c.JSON(http.StatusOK, destination)
}

func (handler *DestinationHandler) UpdateDestination(c *gin.Context) {
	id := c.Param("id")

	var updatedDestination entities.Destination

	if err := c.BindJSON(&updatedDestination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	err := validate.Struct(updatedDestination)

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

	destination, err := handler.Service.UpdateDestination(id, updatedDestination)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "destination not found"})
		return
	}

	c.JSON(http.StatusOK, destination)
}
