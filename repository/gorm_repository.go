package repository

import (
	"fmt"

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

func (repo GormRepository) FindAll(data interface{}) error {
	result := repo.db.Find(data)
	return result.Error
}

func (repo GormRepository) FindByID(data interface{}, id int) error {
	result := repo.db.Find(data, id)
	return result.Error
}

func (repo GormRepository) FindByField(data interface{}, fieldName string, fieldValue interface{}) error {
	result := repo.db.Where(fmt.Sprintf("%s = ?", fieldName), fieldValue).First(data)
	return result.Error
}
