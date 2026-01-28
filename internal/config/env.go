package config

import (
	"go-api-boilerplate/pkg/common/logger"
	"go-api-boilerplate/pkg/validate"
	"os"

	"github.com/joho/godotenv"
)

type Environment string

const (
	Development Environment = "development"
	Staging     Environment = "staging"
	Production  Environment = "production"
	Test        Environment = "test"
)

type Env struct {
	Port        string      `mapstructure:"PORT"`
	Environment Environment `mapstructure:"APP_ENV"     validate:"required,oneof=development staging production test"`

	DBHost     string `mapstructure:"DB_HOST"     validate:"required"`
	DBPort     string `mapstructure:"DB_PORT"     validate:"required"`
	DBDatabase string `mapstructure:"DB_DATABASE" validate:"required"`
	DBUser     string `mapstructure:"DB_USER"     validate:"required"`
	DBPassword string `mapstructure:"DB_PASSWORD" validate:"required"`
	DBSchema   string `mapstructure:"DB_SCHEMA"   validate:"required"`

	GooseDriver       string `mapstructure:"GOOSE_DRIVER"       validate:"required"`
	GooseDBString     string `mapstructure:"GOOSE_DBSTRING"     validate:"required"`
	GooseMigrationDir string `mapstructure:"GOOSE_MIGRATION_DIR" validate:"required"`
	GooseTable        string `mapstructure:"GOOSE_TABLE"        validate:"required"`
}

func InitEnv() *Env {
	env := &Env{}

	if err := env.loadEnv(); err != nil {
		panic(err)
	}
	return env
}

func (env *Env) loadEnv() error {
	_ = godotenv.Load()

	env.Port = os.Getenv("PORT")
	env.Environment = Environment(os.Getenv("APP_ENV"))

	env.DBHost = os.Getenv("DB_HOST")
	env.DBPort = os.Getenv("DB_PORT")
	env.DBDatabase = os.Getenv("DB_DATABASE")
	env.DBUser = os.Getenv("DB_USER")
	env.DBPassword = os.Getenv("DB_PASSWORD")
	env.DBSchema = os.Getenv("DB_SCHEMA")

	env.GooseDriver = os.Getenv("GOOSE_DRIVER")
	env.GooseDBString = os.Getenv("GOOSE_DBSTRING")
	env.GooseMigrationDir = os.Getenv("GOOSE_MIGRATION_DIR")
	env.GooseTable = os.Getenv("GOOSE_TABLE")

	return env.validateEnv()
}

func (e *Env) validateEnv() error {
	if e.Environment == "" {
		e.Environment = Development
	}
	if e.Port == "" {
		e.Port = "8080"
	}

	if err := validate.Validate(e); err != nil {
		logger.Logger.Error("Failed to validate environment variables", "error", err, "env", e)
		return err
	}
	return nil
}
