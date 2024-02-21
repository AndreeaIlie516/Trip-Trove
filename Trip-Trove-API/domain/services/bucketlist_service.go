package services

import (
	"Trip-Trove-API/domain/entities"
	"Trip-Trove-API/domain/repositories"
	"errors"
	"fmt"
)

type BucketListService struct {
	Repo     repositories.BucketListRepository
	UserRepo repositories.UserRepository
}

type BucketListDetails struct {
	BucketList entities.BucketList
	User       entities.User
}

type BucketListsByUser struct {
	User        entities.User
	BucketLists []entities.BucketList
}

func (service *BucketListService) AllBucketLists() ([]entities.BucketList, error) {
	bucketLists, err := service.Repo.AllBucketLists()
	if err != nil {
		return nil, err
	}
	return bucketLists, nil
}

func (service *BucketListService) BucketListByID(idStr string) (*BucketListDetails, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return nil, errors.New("invalid ID format")
	}

	bucketList, err := service.Repo.BucketListByID(id)
	if err != nil {
		return nil, err
	}

	user, err := service.UserRepo.UserByID(bucketList.UserID)
	if err != nil {
		return nil, err
	}

	bucketListDetails := &BucketListDetails{
		BucketList: *bucketList,
		User:       *user,
	}
	return bucketListDetails, nil
}

func (service *BucketListService) BucketListsByUserID(userIDStr string) (*BucketListsByUser, error) {
	var userID uint
	if _, err := fmt.Sscanf(userIDStr, "%d", &userID); err != nil {
		return nil, errors.New("invalid ID format")
	}

	user, err := service.UserRepo.UserByID(userID)
	if err != nil {
		return &BucketListsByUser{}, err
	}

	bucketListsIDs, err := service.Repo.BucketListIDsForUser(userID)
	var bucketLists []entities.BucketList

	for _, bucketListID := range bucketListsIDs {
		bucketList, err := service.Repo.BucketListByID(bucketListID)
		if err != nil {
			return &BucketListsByUser{}, err
		}
		bucketLists = append(bucketLists, *bucketList)
	}

	bucketListsByUser := &BucketListsByUser{
		User:        *user,
		BucketLists: bucketLists,
	}

	return bucketListsByUser, nil
}

func (service *BucketListService) CreateBucketList(bucketList entities.BucketList) (entities.BucketList, error) {
	_, err := service.UserRepo.UserByID(bucketList.UserID)
	if err != nil {
		return entities.BucketList{}, err
	}

	bucketList, err = service.Repo.CreateBucketList(bucketList)
	if err != nil {
		return entities.BucketList{}, err
	}
	return bucketList, nil
}

func (service *BucketListService) DeleteBucketList(idStr string) (entities.BucketList, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.BucketList{}, errors.New("invalid ID format")
	}

	bucketList, err := service.Repo.DeleteBucketList(id)
	if err != nil {
		return entities.BucketList{}, err
	}
	return bucketList, nil
}

func (service *BucketListService) UpdateBucketList(idStr string, bucketList entities.BucketList) (entities.BucketList, error) {
	var id uint
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return entities.BucketList{}, errors.New("invalid ID format")
	}

	bucketList, err := service.Repo.UpdateBucketList(id, bucketList)
	if err != nil {
		return entities.BucketList{}, err
	}
	return bucketList, nil
}
