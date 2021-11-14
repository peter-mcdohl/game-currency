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

func getCurrency(c *gin.Context) {}

func addConvertionRate(c *gin.Context) {}

func convertCurrency(c *gin.Context) {}
