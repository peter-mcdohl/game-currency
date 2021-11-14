package service

import (
	"encoding/json"
	"errors"
	"game-currency/entity"
	"game-currency/repository"
	"log"

	"gorm.io/gorm"
)

type ConversionRateService struct {
	repository repository.Repository
}

func NewConversionRateService(repo repository.Repository) ConversionRateService {
	return ConversionRateService{
		repository: repo,
	}
}

func (crs *ConversionRateService) AddConvertionRate(data repository.GormConversionRate) error {
	// validate entity
	var entity entity.ConversionRate

	b, _ := json.Marshal(data)
	json.Unmarshal(b, &entity)

	if !entity.Valid() {
		err := errors.New("conversion rate data is not valid")
		log.Println(err)
		return err
	}

	// check if both currency exists
	var currencyFrom, currencyTo repository.GormCurrency
	if err := crs.repository.FindByID(&currencyFrom, data.CurrencyIDFrom); err != nil {
		log.Println(err)
		return err
	}
	if err := crs.repository.FindByID(&currencyTo, data.CurrencyIDTo); err != nil {
		log.Println(err)
		return err
	}

	// get conversion rate by field (from & to)
	var conversionRateFrom, conversionRateTo repository.GormConversionRate
	errFrom := crs.repository.FindByField(&conversionRateFrom, "currency_id_from", data.CurrencyIDFrom)
	errTo := crs.repository.FindByField(&conversionRateTo, "currency_id_to", data.CurrencyIDTo)
	if errFrom == nil && errTo == nil {
		err := errors.New("conversion rate is already exists")
		log.Println(err)
		return err
	} else if errors.Is(errFrom, gorm.ErrRecordNotFound) && errors.Is(errTo, gorm.ErrRecordNotFound) {
		errInsert := crs.repository.Insert(&data)
		if errInsert != nil {
			log.Println("failed to add conversion rate", errInsert)
		}

		dataOpposite := repository.GormConversionRate{
			CurrencyIDFrom: data.CurrencyIDTo,
			CurrencyIDTo:   data.CurrencyIDFrom,
			Rate:           1 / data.Rate,
		}
		errInsertOpposite := crs.repository.Insert(&dataOpposite)
		if errInsertOpposite != nil {
			log.Println("failed to add opposite conversion rate", errInsertOpposite)
		}

		if errInsert != nil && errInsertOpposite != nil {
			return errors.New("failed to add conversion rate")
		}
	} else {
		log.Println("conversion rate from", errFrom)
		log.Println("conversion rate to", errTo)
		return errors.New("failed to add conversion rate")
	}

	return nil
}

type ConversionRateRequest struct {
	CurrencyIDFrom int     `json:"currency_id_from"`
	CurrencyIDTo   int     `json:"currency_id_to"`
	Amount         float64 `json:"amount"`
}

func (crs *ConversionRateService) ConvertCurrency(data ConversionRateRequest) (float64, error) {
	var conversionRate repository.GormConversionRate

	if err := crs.repository.FindByConditionStruct(&conversionRate,
		&repository.GormConversionRate{
			CurrencyIDFrom: data.CurrencyIDFrom,
			CurrencyIDTo:   data.CurrencyIDTo,
		},
	); err != nil {
		log.Println(err)
		return 0, err
	}

	b, _ := json.Marshal(conversionRate)
	var entity entity.ConversionRate
	json.Unmarshal(b, &entity)

	return entity.Convert(data.Amount), nil
}
