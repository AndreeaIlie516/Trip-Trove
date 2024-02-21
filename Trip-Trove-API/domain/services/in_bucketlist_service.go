package services

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/domain/repositories"
	"errors"
	"fmt"
)

type InBucketListService struct {
	Repo            repositories.InBucketListRepository
	DestinationRepo repositories.DestinationRepository
	BucketListRepo  repositories.BucketListRepository
}

type InBucketListDetail struct {
	Association entities.InBucketList
	Destination entities.Destination
	BucketList  entities.BucketList
}

type DestinationWithBucketLists struct {
	Destination entities.Destination
	BucketLists []entities.BucketList
}

type BucketListWithDestinations struct {
	BucketList   entities.BucketList
	Destinations []entities.Destination
}

func (service *InBucketListService) AllInBucketListAssociations() ([]entities.InBucketList, error) {
	inBucketListAssociations, err := service.Repo.AllInBucketListAssociations()
	if err != nil {
		return nil, err
	}
	return inBucketListAssociations, nil
}

func (service *InBucketListService) InBucketListAssociationByID(idStr string) (*InBucketListDetail, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return nil, errors.New("invalid ID format")
	}

	inBucketListAssociation, err := service.Repo.InBucketListAssociationByID(id)
	if err != nil {
		return nil, err
	}

	destination, err := service.DestinationRepo.DestinationByID(inBucketListAssociation.DestinationID)
	if err != nil {
		return nil, err
	}

	bucketList, err := service.BucketListRepo.BucketListByID(inBucketListAssociation.BucketListID)
	if err != nil {
		return nil, err
	}

	inBucketListDetail := &InBucketListDetail{
		Association: *inBucketListAssociation,
		Destination: *destination,
		BucketList:  *bucketList,
	}
	return inBucketListDetail, nil
}

func (service *InBucketListService) InBucketListAssociation(destinationIDStr string, bucketListIDStr string) (*InBucketListDetail, error) {
	var destinationID uint
	if _, err := fmt.Sscanf(destinationIDStr, "%d", &destinationID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	destination, err := service.DestinationRepo.DestinationByID(destinationID)
	if err != nil {
		return nil, err
	}

	var bucketListID uint
	if _, err := fmt.Sscanf(bucketListIDStr, "%d", &bucketListID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	bucketList, err := service.BucketListRepo.BucketListByID(bucketListID)
	if err != nil {
		return nil, err
	}

	inBucketListAssociation, err := service.Repo.InBucketListAssociation(destinationID, bucketListID)
	if err != nil {
		return nil, err
	}

	inBucketListDetail := &InBucketListDetail{
		Association: *inBucketListAssociation,
		Destination: *destination,
		BucketList:  *bucketList,
	}
	return inBucketListDetail, nil
}

func (service *InBucketListService) DestinationWithBucketLists(destinationIDStr string) (*DestinationWithBucketLists, error) {
	var destinationID uint
	if _, err := fmt.Sscanf(destinationIDStr, "%d", &destinationID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	destination, err := service.DestinationRepo.DestinationByID(destinationID)
	if err != nil {
		return &DestinationWithBucketLists{}, err
	}

	bucketListIDs, err := service.Repo.BucketListIDsForDestination(destinationID)
	var bucketLists []entities.BucketList

	for _, bucketListID := range bucketListIDs {
		bucketList, err := service.BucketListRepo.BucketListByID(bucketListID)
		if err != nil {
			return &DestinationWithBucketLists{}, err
		}
		bucketLists = append(bucketLists, *bucketList)
	}

	DestinationWithBucketList := &DestinationWithBucketLists{
		Destination: *destination,
		BucketLists: bucketLists,
	}

	return DestinationWithBucketList, nil
}

func (service *InBucketListService) BucketListWithDestinations(bucketListIDStr string) (*BucketListWithDestinations, error) {
	var bucketListID uint
	if _, err := fmt.Sscanf(bucketListIDStr, "%d", &bucketListID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	bucketList, err := service.BucketListRepo.BucketListByID(bucketListID)
	if err != nil {
		return &BucketListWithDestinations{}, err
	}

	destinationIDs, err := service.Repo.DestinationIDsForBucketList(bucketListID)
	var destinations []entities.Destination

	for _, destinationID := range destinationIDs {
		destination, err := service.DestinationRepo.DestinationByID(destinationID)
		if err != nil {
			return &BucketListWithDestinations{}, err
		}
		destinations = append(destinations, *destination)
	}

	bucketListWithDestinations := &BucketListWithDestinations{
		BucketList:   *bucketList,
		Destinations: destinations,
	}

	return bucketListWithDestinations, nil
}

func (service *InBucketListService) AddDestinationToBucketList(inBucketListAssociation entities.InBucketList) (entities.InBucketList, error) {
	_, err := service.DestinationRepo.DestinationByID(inBucketListAssociation.DestinationID)
	if err != nil {
		return entities.InBucketList{}, err
	}

	_, err = service.BucketListRepo.BucketListByID(inBucketListAssociation.BucketListID)
	if err != nil {
		return entities.InBucketList{}, err
	}

	inBucketListAssociation, err = service.Repo.AddDestinationToBucketList(inBucketListAssociation)
	if err != nil {
		return entities.InBucketList{}, err
	}
	return inBucketListAssociation, nil
}

func (service *InBucketListService) DeleteDestinationFromBucketList(idStr string) (entities.InBucketList, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.InBucketList{}, errors.New("invalid ID format")
	}

	inBucketListAssociation, err := service.Repo.DeleteDestinationFromBucketList(id)
	if err != nil {
		return entities.InBucketList{}, err
	}
	return inBucketListAssociation, nil
}

func (service *InBucketListService) DeleteDestinationFromBucketLists(destinationIDStr string) ([]entities.InBucketList, error) {
	var destinationID uint
	if _, err := fmt.Sscanf(destinationIDStr, "%d", &destinationID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	_, err := service.DestinationRepo.DestinationByID(destinationID)
	if err != nil {
		return []entities.InBucketList{}, err
	}

	inBucketListAssociation, err := service.Repo.DeleteDestinationFromItsBucketLists(destinationID)
	if err != nil {
		return []entities.InBucketList{}, err
	}
	return inBucketListAssociation, nil
}

func (service *InBucketListService) DeleteBucketListFromItsDestinations(bucketListIDStr string) ([]entities.InBucketList, error) {
	var bucketListID uint
	if _, err := fmt.Sscanf(bucketListIDStr, "%d", &bucketListID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	_, err := service.BucketListRepo.BucketListByID(bucketListID)
	if err != nil {
		return []entities.InBucketList{}, err
	}

	inBucketListAssociation, err := service.Repo.DeleteBucketListFromItsDestinations(bucketListID)
	if err != nil {
		return []entities.InBucketList{}, err
	}
	return inBucketListAssociation, nil
}
