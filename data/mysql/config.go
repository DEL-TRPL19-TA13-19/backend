package mysql

import (
	"fmt"
	"github.com/spf13/viper"
)

type databaseConfiguration struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	DBType     string
}

func readDatabaseConfiguration() databaseConfiguration {
	var config databaseConfiguration

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	config.DBUser = viper.GetString("DBUser")
	config.DBPassword = viper.GetString("DBPassword")
	config.DBName = viper.GetString("DBName")
	config.DBHost = viper.GetString("DBHost")
	config.DBPort = viper.GetString("DBPort")
	config.DBType = viper.GetString("DBType")

	return config
}

func GetDBType() string {
	config := readDatabaseConfiguration()
	dbType := fmt.Sprintf("%s", config.DBType)

	return dbType
}

func GetMySQLConnectionString() string {
	config := readDatabaseConfiguration()

	db := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName)

	return db
}
