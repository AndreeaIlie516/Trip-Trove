package main

import (
	"Trip-Trove-API/database"
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/domain/services"
	"Trip-Trove-API/infrastructure/dataaccess"
	"Trip-Trove-API/infrastructure/middlewares"
	"Trip-Trove-API/presentation/handlers"
	"Trip-Trove-API/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	router := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.ConnectDB()

	entitiesToMigrate := []interface{}{
		&entities.User{},
	}

	for _, entity := range entitiesToMigrate {
		err := db.AutoMigrate(entity)
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
	}

	authMiddleware := middlewares.AuthMiddleware{}

	userRepository := dataaccess.NewGormUserRepository(db)

	userService := services.UserService{Repo: userRepository}

	userHandler := handlers.UserHandler{Service: &userService}

	routes.RegisterUserRoutes(router, &userHandler, authMiddleware)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}
