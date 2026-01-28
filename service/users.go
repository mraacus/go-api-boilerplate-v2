package service

import (
	"go-api-boilerplate/internal/dao"
	"go-api-boilerplate/pkg/common/logger"
)

func CreateUser(name string, role string) error {
	user := &dao.User{
		Name: name,
		Role: role,
	}
	err := dao.CreateUser(user)
	if err != nil {
		logger.Logger.Error("Failed to create user", "error", err)
		return err
	}
	return nil
}

func ListUsers() ([]dao.User, error) {
	users, err := dao.ListAllUsers()
	if err != nil {
		logger.Logger.Error("Failed to list users", "error", err)
		return nil, err
	}
	return users, nil
}