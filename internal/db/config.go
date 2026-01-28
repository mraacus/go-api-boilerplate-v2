package db

import (
	"go-api-boilerplate/pkg/common/logger"
	"os"
)

type DBConfig struct {
	DatabaseName string `mapstructure:"DB_DATABASE"`
	Password     string `mapstructure:"DB_PASSWORD"`
	Username     string `mapstructure:"DB_USER"`
	Port         string `mapstructure:"DB_PORT"`
	Host         string `mapstructure:"DB_HOST"`
	Schema       string `mapstructure:"DB_SCHEMA"`
}

func LoadConfig() DBConfig {
	database_name := os.Getenv("DB_DATABASE")
	password      := os.Getenv("DB_PASSWORD")
	username      := os.Getenv("DB_USER")
	port          := os.Getenv("DB_PORT")
	host          := os.Getenv("DB_HOST")
	schema        := os.Getenv("DB_SCHEMA")

	if database_name == "" || password == "" || username == "" || port == "" || host == "" || schema == "" {
		logger.Logger.Error("Database configuration is incomplete", 
            "host", host, 
            "username", username, 
            "database", database_name, 
            "port", port,
            "schema", schema)
        panic("incomplete database configuration")
	}

	return DBConfig{
		DatabaseName: database_name,
		Password:     password,
		Username:     username,
		Port:         port,
		Host:         host,
		Schema:      schema,
	}
}
