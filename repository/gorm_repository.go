package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewPostgresDB(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), opts...)
}

func NewGormRepository(db *gorm.DB) GormRepository {
	return GormRepository{
		db: db,
	}
}

func (repo GormRepository) Insert(data interface{}) error {
	result := repo.db.Create(data)
	return result.Error
}
