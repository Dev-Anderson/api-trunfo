package database

import (
	"api-trunfo/config"
	"api-trunfo/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func GetDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.LoadEnv().Host, config.LoadEnv().User, config.LoadEnv().Password, config.LoadEnv().Database)
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalln("Falha ao conectar com o banco de dados")
	}

	DB.AutoMigrate(&models.Veiculos{}, &models.Login{}, &models.User{})

}
