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
		&entities.InBucketList{},
		&entities.BucketList{},
		&entities.Destination{},
		&entities.Location{},
		&entities.User{},
	}

	for _, entity := range entitiesToMigrate {
		err := db.AutoMigrate(entity)
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
	}

	authMiddleware := middlewares.AuthMiddleware{}

	inBucketListRepository := dataaccess.NewGormInBucketListRepository(db)
	bucketListRepository := dataaccess.NewGormBucketListRepository(db)
	destinationRepository := dataaccess.NewGormDestinationRepository(db)
	locationRepository := dataaccess.NewGormLocationRepository(db)
	userRepository := dataaccess.NewGormUserRepository(db)

	inBucketListService := services.InBucketListService{Repo: inBucketListRepository, DestinationRepo: destinationRepository, BucketListRepo: bucketListRepository}
	bucketListService := services.BucketListService{Repo: bucketListRepository, UserRepo: userRepository}
	destinationService := services.DestinationService{Repo: destinationRepository, LocationRepo: locationRepository}
	locationService := services.LocationService{Repo: locationRepository}
	userService := services.UserService{Repo: userRepository}

	inBucketListHandler := handlers.InBucketListHandler{Service: &inBucketListService}
	bucketListHandler := handlers.BucketListHandler{Service: &bucketListService}
	destinationHandler := handlers.DestinationHandler{Service: &destinationService}
	locationHandler := handlers.LocationHandler{Service: &locationService}
	userHandler := handlers.UserHandler{Service: &userService}

	routes.RegisterInBucketListRoutes(router, &inBucketListHandler)
	routes.RegisterBucketListRoutes(router, &bucketListHandler)
	routes.RegisterDestinationRoutes(router, &destinationHandler)
	routes.RegisterLocationRoutes(router, &locationHandler)
	routes.RegisterUserRoutes(router, &userHandler, authMiddleware)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}
