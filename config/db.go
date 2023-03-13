package config

import "fmt"

const (
	DBUser     = "user"
	DBPassword = "password"
	DBName     = "ta13db"
	DBHost     = "localhost"
	DBPort     = "3306"
	DBType     = "mysql"
	//DBTimeZone = "Asia/Jakarta"
)

func GetDBType() string {
	return DBType
}

func GetMySQLConnectionString() string {
	db := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBHost,
		DBPort,
		DBName)

	return db
}
