package repositories

import "Trip-Trove-API/domain/entities"

type InBucketListRepository interface {
	AllInBucketListAssociations() ([]entities.InBucketList, error)
	AllInBucketListAssociationIDs() ([]uint, error)
	InBucketListAssociationByID(id uint) (*entities.InBucketList, error)
	InBucketListAssociation(destinationID uint, bucketListID uint) (*entities.InBucketList, error)
	BucketListIDsForDestination(destinationID uint) ([]uint, error)
	DestinationIDsForBucketList(bucketListID uint) ([]uint, error)
	AddDestinationToBucketList(finBucketListAssociation entities.InBucketList) (entities.InBucketList, error)
	DeleteDestinationFromBucketList(id uint) (entities.InBucketList, error)
	DeleteDestinationFromItsBucketLists(destinationID uint) ([]entities.InBucketList, error)
	DeleteBucketListFromItsDestinations(bucketListID uint) ([]entities.InBucketList, error)
}
