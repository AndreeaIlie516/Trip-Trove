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
		&entities.BucketList{},
		&entities.City{},
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

	bucketListRepository := dataaccess.NewGormBucketListRepository(db)
	cityRepository := dataaccess.NewGormCityRepository(db)
	destinationRepository := dataaccess.NewGormDestinationRepository(db)
	locationRepository := dataaccess.NewGormLocationRepository(db)
	userRepository := dataaccess.NewGormUserRepository(db)

	bucketListService := services.BucketListService{Repo: bucketListRepository, UserRepo: userRepository}
	cityService := services.CityService{Repo: cityRepository, LocationRepo: locationRepository}
	destinationService := services.DestinationService{Repo: destinationRepository, LocationRepo: locationRepository, CityRepo: cityRepository}
	locationService := services.LocationService{Repo: locationRepository, CityRepo: cityRepository}
	userService := services.UserService{Repo: userRepository}

	bucketListHandler := handlers.BucketListHandler{Service: &bucketListService}
	cityHandler := handlers.CityHandler{Service: &cityService}
	destinationHandler := handlers.DestinationHandler{Service: &destinationService}
	locationHandler := handlers.LocationHandler{Service: &locationService}
	userHandler := handlers.UserHandler{Service: &userService}

	routes.RegisterBucketListRoutes(router, &bucketListHandler)
	routes.RegisterCityRoutes(router, &cityHandler)
	routes.RegisterDestinationRoutes(router, &destinationHandler)
	routes.RegisterLocationRoutes(router, &locationHandler)
	routes.RegisterUserRoutes(router, &userHandler, authMiddleware)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}
