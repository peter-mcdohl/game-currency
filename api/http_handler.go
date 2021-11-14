package api

import (
	"game-currency/repository"
	"game-currency/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addCurrency(c *gin.Context) {
	var data repository.GormCurrency
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
		return
	}

	repo := repository.NewGormRepository(gormDB)
	svc := service.NewCurrencyService(repo)

	if err := svc.AddCurrency(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to add currency"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Succuessfully add currency"})
}

func getCurrency(c *gin.Context) {
	repo := repository.NewGormRepository(gormDB)
	svc := service.NewCurrencyService(repo)

	data := svc.GetAllCurrency()

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func addConvertionRate(c *gin.Context) {
	var data repository.GormConversionRate
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
		return
	}

	repo := repository.NewGormRepository(gormDB)
	svc := service.NewConversionRateService(repo)

	if err := svc.AddConvertionRate(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to add conversion rate"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Succuessfully add conversion rate"})
}

func convertCurrency(c *gin.Context) {
	var data service.ConversionRateRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
		return
	}

	repo := repository.NewGormRepository(gormDB)
	svc := service.NewConversionRateService(repo)

	result, err := svc.ConvertCurrency(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to convert"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
