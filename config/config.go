package config

import (
	"database/sql"
	"log"
)

type Config struct {
	DB *sql.DB
}

func InitConfig() Config {
	dbConn := ConnDB()
	log.Println("Db Connections is initialized")
	return Config{DB: dbConn}
}
