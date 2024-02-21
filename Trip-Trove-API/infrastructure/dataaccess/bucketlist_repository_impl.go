package dataaccess

import (
	"Trip-Trove-API/domain/entities"
	"errors"
	"gorm.io/gorm"
)

type GormBucketListRepository struct {
	Db *gorm.DB
}

func NewGormBucketListRepository(db *gorm.DB) *GormBucketListRepository {
	return &GormBucketListRepository{Db: db}
}

func (r *GormBucketListRepository) AllBucketLists() ([]entities.BucketList, error) {
	var bucketLists []entities.BucketList
	result := r.Db.Find(&bucketLists)
	return bucketLists, result.Error
}

func (r *GormBucketListRepository) AllBucketListIDs() ([]uint, error) {
	var bucketListIDs []uint

	if err := r.Db.Model(&entities.BucketList{}).Select("ID").Find(&bucketListIDs).Error; err != nil {
		return nil, err
	}

	return bucketListIDs, nil
}

func (r *GormBucketListRepository) BucketListByID(id uint) (*entities.BucketList, error) {
	var bucketList entities.BucketList

	if err := r.Db.First(&bucketList, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("bucketList not found")
		}
		return nil, err
	}

	return &bucketList, nil
}

func (r *GormBucketListRepository) BucketListIDsForUser(userID uint) ([]uint, error) {
	var bucketListIDs []uint

	if err := r.Db.Where("city_id = ?", userID).Model(&entities.BucketList{}).Select("ID").Find(&bucketListIDs).Error; err != nil {
		return nil, err
	}

	return bucketListIDs, nil
}

func (r *GormBucketListRepository) CreateBucketList(bucketList entities.BucketList) (entities.BucketList, error) {
	if err := r.Db.Create(&bucketList).Error; err != nil {
		return entities.BucketList{}, err
	}
	return bucketList, nil
}

func (r *GormBucketListRepository) DeleteBucketList(id uint) (entities.BucketList, error) {
	var bucketList entities.BucketList

	if err := r.Db.First(&bucketList, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.BucketList{}, errors.New("bucketList not found")
		}
		return entities.BucketList{}, err
	}

	if err := r.Db.Delete(&bucketList).Error; err != nil {
		return entities.BucketList{}, err
	}

	return bucketList, nil
}

func (r *GormBucketListRepository) UpdateBucketList(id uint, updatedBucketList entities.BucketList) (entities.BucketList, error) {
	var bucketList entities.BucketList

	if err := r.Db.First(&bucketList, "ID = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.BucketList{}, errors.New("bucketList not found")
		}
		return entities.BucketList{}, err
	}

	if err := r.Db.Model(&bucketList).Updates(updatedBucketList).Error; err != nil {
		return entities.BucketList{}, err
	}

	return bucketList, nil
}
