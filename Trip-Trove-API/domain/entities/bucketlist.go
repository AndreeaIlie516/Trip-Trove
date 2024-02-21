package entities

import "gorm.io/gorm"

type BucketList struct {
	gorm.Model
	UserID      uint   `gorm:"column:user_id;not null" json:"user_id" validate:"required,number"`
	Name        string `gorm:"column:name;not null" json:"name" validate:"required,min=3,max=30"`
	Description string `gorm:"column:description" json:"description" validate:"max=256"`
}
