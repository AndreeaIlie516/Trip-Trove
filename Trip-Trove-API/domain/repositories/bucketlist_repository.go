package repositories

import "Trip-Trove-API/domain/entities"

type BucketListRepository interface {
	AllBucketLists() ([]entities.BucketList, error)
	AllBucketListIDs() ([]uint, error)
	BucketListByID(id uint) (*entities.BucketList, error)
	BucketListIDsForUser(userID uint) ([]uint, error)
	CreateBucketList(bucketList entities.BucketList) (entities.BucketList, error)
	UpdateBucketList(id uint, updatedBucketList entities.BucketList) (entities.BucketList, error)
	DeleteBucketList(id uint) (entities.BucketList, error)
}
