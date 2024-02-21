package dataaccess

import (
	"Trip-Trove-API/domain/entities"
	"errors"
	"gorm.io/gorm"
)

type GormInBucketListRepository struct {
	Db *gorm.DB
}

func NewGormInBucketListRepository(db *gorm.DB) *GormInBucketListRepository {
	return &GormInBucketListRepository{Db: db}
}

func (r *GormInBucketListRepository) AllInBucketListAssociations() ([]entities.InBucketList, error) {
	var inBucketListAssociations []entities.InBucketList
	result := r.Db.Find(&inBucketListAssociations)
	return inBucketListAssociations, result.Error
}

func (r *GormInBucketListRepository) AllInBucketListAssociationIDs() ([]uint, error) {
	var inBucketListAssociationIDs []uint

	if err := r.Db.Model(&entities.InBucketList{}).Select("ID").Find(&inBucketListAssociationIDs).Error; err != nil {
		return nil, err
	}

	return inBucketListAssociationIDs, nil
}

func (r *GormInBucketListRepository) InBucketListAssociationByID(id uint) (*entities.InBucketList, error) {
	var inBucketListAssociation entities.InBucketList

	if err := r.Db.First(&inBucketListAssociation, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("inBucketListAssociation not found")
		}
		return nil, err
	}

	return &inBucketListAssociation, nil
}

func (r *GormInBucketListRepository) InBucketListAssociation(destinationID uint, bucketListID uint) (*entities.InBucketList, error) {
	var inBucketListAssociation entities.InBucketList

	if err := r.Db.First(&inBucketListAssociation, "destination_id = ? AND bucket_list_id = ?", destinationID, bucketListID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("inBucketListAssociation not found")
		}
		return nil, err
	}

	return &inBucketListAssociation, nil
}

func (r *GormInBucketListRepository) BucketListIDsForDestination(destinationID uint) ([]uint, error) {
	var bucketListIDs []uint

	if err := r.Db.Where("destination_id = ?", destinationID).Model(&entities.InBucketList{}).Select("ID").Find(&bucketListIDs).Error; err != nil {
		return nil, err
	}

	return bucketListIDs, nil
}

func (r *GormInBucketListRepository) DestinationIDsForBucketList(bucketListID uint) ([]uint, error) {
	var destinationIDs []uint

	if err := r.Db.Where("bucket_list_id = ?", bucketListID).Model(&entities.InBucketList{}).Select("destination_id").Find(&destinationIDs).Error; err != nil {
		return nil, err
	}

	return destinationIDs, nil
}

func (r *GormInBucketListRepository) AddDestinationToBucketList(inBucketListAssociation entities.InBucketList) (entities.InBucketList, error) {
	if err := r.Db.Create(&inBucketListAssociation).Error; err != nil {
		return entities.InBucketList{}, err
	}
	return inBucketListAssociation, nil
}

func (r *GormInBucketListRepository) DeleteDestinationFromBucketList(id uint) (entities.InBucketList, error) {
	var inBucketListAssociation entities.InBucketList

	if err := r.Db.First(&inBucketListAssociation, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.InBucketList{}, errors.New("inBucketListAssociation not found")
		}
		return entities.InBucketList{}, err
	}

	if err := r.Db.Delete(&inBucketListAssociation).Error; err != nil {
		return entities.InBucketList{}, err
	}

	return inBucketListAssociation, nil
}

func (r *GormInBucketListRepository) DeleteDestinationFromItsBucketLists(destinationID uint) ([]entities.InBucketList, error) {
	var inBucketListAssociations []entities.InBucketList
	var deletedAssociations []entities.InBucketList

	if err := r.Db.Where("destination_id = ?", destinationID).Find(&inBucketListAssociations).Error; err != nil {
		return []entities.InBucketList{}, err
	}

	deletedAssociations = append(deletedAssociations, inBucketListAssociations...)

	if err := r.Db.Delete(&inBucketListAssociations, "destination_id = ?", destinationID).Error; err != nil {
		return []entities.InBucketList{}, err
	}

	return deletedAssociations, nil
}

func (r *GormInBucketListRepository) DeleteBucketListFromItsDestinations(bucketListID uint) ([]entities.InBucketList, error) {
	var inBucketListAssociations []entities.InBucketList
	var deletedAssociations []entities.InBucketList

	if err := r.Db.Where("bucket_list_id = ?", bucketListID).Find(&inBucketListAssociations).Error; err != nil {
		return []entities.InBucketList{}, err
	}

	deletedAssociations = append(deletedAssociations, inBucketListAssociations...)

	if err := r.Db.Delete(&inBucketListAssociations, "bucket_list_id = ?", bucketListID).Error; err != nil {
		return []entities.InBucketList{}, err
	}

	return deletedAssociations, nil
}
