package db

import (
	"fmt"
	"go-api-boilerplate/pkg/common/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var (
	readDB  *gorm.DB
	writeDB *gorm.DB
)

func InitDB() {
	config := LoadConfig()
	InitClient(config)
}

func InitClient(config DBConfig) {
	dsn := getDsn(config)

	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
	}), &gorm.Config{})
	if err != nil {
		logger.Logger.Error("Failed to connect to database", "error", err, "dsn", dsn)
		panic(err)
	}

	readDB = conn.Clauses(dbresolver.Read)
	writeDB = conn.Clauses(dbresolver.Write)
}

func getDsn(config DBConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", 
			config.Host, 
			config.Username,  
			config.Password, 
			config.DatabaseName, 
			config.Port)
}

func GetReadDB() *gorm.DB {
	return readDB
}

func GetWriteDB() *gorm.DB {
	return writeDB
}