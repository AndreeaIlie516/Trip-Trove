package entities

import "gorm.io/gorm"

type InBucketList struct {
	gorm.Model
	DestinationID uint   `gorm:"column:destination_id;not null" json:"destination_id" validate:"required,number"`
	BucketListID  uint   `gorm:"column:bucket_list_id;not null" json:"bucket_list_id" validate:"required,number"`
	CheckInDate   string `gorm:"column:check_in_date" json:"check_in_date" validate:"max=20"`
	CheckOutDate  string `gorm:"column:check_out_date" json:"check_out_date" validate:"max=20"`
}
