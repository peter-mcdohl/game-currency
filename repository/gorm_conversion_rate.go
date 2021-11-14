package repository

import (
	"time"

	"gorm.io/gorm"
)

type GormConversionRate struct {
	ID             uint           `json:"id" gorm:"primarykey"`
	CurrencyIDFrom int            `json:"currency_id_from" gorm:"index:,column:currency_id_from"`
	CurrencyIDTo   int            `json:"currency_id_to" gorm:"index:,column:currency_id_to"`
	Rate           float64        `json:"currency_rate"`
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

func (GormConversionRate) TableName() string {
	return "conversion_rate"
}
