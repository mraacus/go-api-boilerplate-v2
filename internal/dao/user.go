package dao

import (
	"go-api-boilerplate/internal/db"
	"go-api-boilerplate/pkg/common/logger"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const (
	tableName = "users"
)

type User struct {
	ID        uint 	`gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Role      string 
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

func CreateUser(user *User) error {
	ctxMsg := "user.CreateUser"

	q := db.GetWriteDB().Table(tableName)

	res := q.Create(user)
	if res.Error != nil {
		logger.Logger.Error("Failed to create user", "error", res.Error, "user", user)
		return errors.Wrap(res.Error, ctxMsg)
	}
	logger.Logger.Info("Successfully created user", "user", user)

	return nil
}

func ListAllUsers() (users []User, err error) {
	ctxMsg := "user.ListAllUsers"

	q := db.GetReadDB().Table(tableName)

	res := q.Find(&users)
	if res.Error != nil {
		logger.Logger.Error("Failed to list users", "error", res.Error)
		return nil, errors.Wrap(res.Error, ctxMsg)
	}

	return users, nil
}