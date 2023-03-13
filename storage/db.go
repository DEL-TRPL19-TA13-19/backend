package storage

import (
	"github.com/jinzhu/gorm"
	"log"
	"ta13-svc/config"
	"ta13-svc/model"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	connectionString := config.GetMySQLConnectionString()
	log.Print(connectionString)

	var err error

	DB, err = gorm.Open(config.GetDBType(), connectionString)

	if err != nil {
		log.Panic(err)
	}

	DB.AutoMigrate(model.User{})

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
