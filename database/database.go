package database

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"strings"
)

var (
	dbConnections map[string]*gorm.DB
)

func Init() {

	dbConfigurations := map[string]Db{
		"TA13DB": &dbMySQL{
			db: db{
				Host: os.Getenv("DB_HOST_TA13DB"),
				User: os.Getenv("DB_USER_TA13DB"),
				Pass: os.Getenv("DB_PASS_TA13DB"),
				Port: os.Getenv("DB_PORT_TA13DB"),
				Name: os.Getenv("DB_NAME_TA13DB"),
			},
			Charset:   "utf8mb4",
			ParseTime: "True",
			Loc:       "Local",
		},
	}

	dbConnections = make(map[string]*gorm.DB)

	for k, v := range dbConfigurations {
		db, err := v.Init()
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to database %s", k))
		}
		dbConnections[k] = db
		logrus.Info(fmt.Sprintf("Successfully connected to database %s", k))
	}
}

func Connection(name string) (*gorm.DB, error) {
	if dbConnections[strings.ToUpper(name)] == nil {
		return nil, errors.New("Connections is undefined")
	}
	return dbConnections[strings.ToUpper(name)], nil
}
