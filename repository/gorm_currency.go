package repository

import (
	"time"

	"gorm.io/gorm"
)

type GormCurrency struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (GormCurrency) TableName() string {
	return "currency"
}
