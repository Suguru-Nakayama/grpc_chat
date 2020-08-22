package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DBMS       = "mysql"
	MYSQL_PORT = 3306
	DB_NAME    = "grpc_chat"
	CHARSET    = "utf8mb4"
	PARSE_TIME = "True"
	LOCATION   = "Local"
)

func GetDBConnection() (*gorm.DB, error) {
	config := fmt.Sprintf("%s:%s@tcp(db:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		MYSQL_PORT,
		DB_NAME,
		CHARSET,
		PARSE_TIME,
		LOCATION,
	)

	db, err := gorm.Open(DBMS, config)
	if err != nil {
		return nil, err
	}

	return db, nil
}
