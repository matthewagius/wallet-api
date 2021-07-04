package util

import (
	"fmt"

	"github.com/matthewagius/wallet-api/config"
)

type DBConfiguration struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}

func NewDatabaseConfig() *DBConfiguration {
	database := &DBConfiguration{
		Port:     config.DB_PORT,
		Host:     config.DB_HOST,
		DBName:   config.DB_DATABASE,
		User:     config.DB_USER,
		Password: config.DB_PASSWORD,
	}
	return database
}

func DBUrlConnection(dbConfig *DBConfiguration) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
