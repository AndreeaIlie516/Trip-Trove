package services

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/domain/repositories"
	"errors"
	"fmt"
)

type DestinationService struct {
	Repo         repositories.DestinationRepository
	LocationRepo repositories.LocationRepository
}

type DestinationDetails struct {
	Destination entities.Destination
	Location    entities.Location
}

type DestinationWithLocation struct {
	Destination entities.Destination
	Location    entities.Location
}

type DestinationsByLocation struct {
	Location     entities.Location
	Destinations []entities.Destination
}

func (service *DestinationService) AllDestinations() ([]entities.Destination, error) {
	destinations, err := service.Repo.AllDestinations()
	if err != nil {
		return nil, err
	}
	return destinations, nil
}

func (service *DestinationService) DestinationByID(idStr string) (*DestinationDetails, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return nil, errors.New("invalid ID format")
	}

	destination, err := service.Repo.DestinationByID(id)
	if err != nil {
		return nil, err
	}

	location, err := service.LocationRepo.LocationByID(destination.LocationID)
	if err != nil {
		return nil, err
	}

	destinationDetails := &DestinationDetails{
		Destination: *destination,
		Location:    *location,
	}
	return destinationDetails, nil
}

func (service *DestinationService) DestinationsByLocationID(locationIDStr string) (*DestinationsByLocation, error) {
	var locationID uint
	if _, err := fmt.Sscanf(locationIDStr, "%d", &locationID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	location, err := service.LocationRepo.LocationByID(locationID)
	if err != nil {
		return &DestinationsByLocation{}, err
	}

	destinationIDs, err := service.Repo.DestinationIDsForLocation(locationID)
	var destinations []entities.Destination

	for _, destinationID := range destinationIDs {
		destination, err := service.Repo.DestinationByID(destinationID)
		if err != nil {
			return &DestinationsByLocation{}, err
		}
		destinations = append(destinations, *destination)
	}

	destinationsByLocation := &DestinationsByLocation{
		Location:     *location,
		Destinations: destinations,
	}

	return destinationsByLocation, nil
}

func (service *DestinationService) CreateDestination(destination entities.Destination) (entities.Destination, error) {
	_, err := service.LocationRepo.LocationByID(destination.LocationID)
	if err != nil {
		return entities.Destination{}, err
	}

	destination, err = service.Repo.CreateDestination(destination)
	if err != nil {
		return entities.Destination{}, err
	}
	return destination, nil
}

func (service *DestinationService) DeleteDestination(idStr string) (entities.Destination, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.Destination{}, errors.New("invalid ID format")
	}

	destination, err := service.Repo.DeleteDestination(id)
	if err != nil {
		return entities.Destination{}, err
	}
	return destination, nil
}

func (service *DestinationService) UpdateDestination(idStr string, destination entities.Destination) (entities.Destination, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.Destination{}, errors.New("invalid ID format")
	}

	destination, err := service.Repo.UpdateDestination(id, destination)
	if err != nil {
		return entities.Destination{}, err
	}
	return destination, nil
}
