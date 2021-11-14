package service

import (
	"encoding/json"
	"errors"
	"game-currency/entity"
	"game-currency/repository"
	"log"

	"gorm.io/gorm"
)

type CurrencyService struct {
	repository repository.Repository
}

func NewCurrencyService(repo repository.Repository) CurrencyService {
	return CurrencyService{
		repository: repo,
	}
}

func (cs *CurrencyService) AddCurrency(data repository.GormCurrency) error {
	var entity entity.Currency

	b, _ := json.Marshal(data)
	json.Unmarshal(b, &entity)

	if !entity.Valid() {
		err := errors.New("currency data is not valid")
		log.Println(err)
		return err
	}

	bEntity, _ := json.Marshal(entity)
	json.Unmarshal(bEntity, &data)

	err := cs.repository.Insert(&data)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (cs *CurrencyService) GetAllCurrency() []repository.GormCurrency {
	var data []repository.GormCurrency

	err := cs.repository.FindAll(&data)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)
	}

	return data
}
