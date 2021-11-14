package api

import (
	"game-currency/repository"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	gormDB *gorm.DB
)

func Run() {
	var errDB error
	// TODO: get it from env var
	dsn := "host=localhost user=postgres password=secret dbname=currency sslmode=disable TimeZone=Asia/Jakarta"
	gormDB, errDB = repository.NewPostgresDB(dsn)
	if errDB != nil {
		log.Fatal(errDB)
	}

	// DB Migration
	if err := gormDB.AutoMigrate(
		&repository.GormCurrency{},
		&repository.GormConversionRate{},
	); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, nil)
	})

	v1 := router.Group("/v1")
	{
		v1.POST("/currency", addCurrency)
		v1.GET("/currency", getCurrency)
		v1.POST("/currency/conversion-rate", addConvertionRate)
	}

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
