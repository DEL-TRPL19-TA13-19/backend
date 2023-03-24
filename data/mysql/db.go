package mysql

import (
	"github.com/jinzhu/gorm"
	"log"
	"ta13-svc/entity"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	connectionString := GetMySQLConnectionString()
	log.Print(connectionString)

	var err error

	DB, err = gorm.Open(GetDBType(), connectionString)

	if err != nil {
		log.Panic(err)
	}

	DB.AutoMigrate(entity.User{})

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
