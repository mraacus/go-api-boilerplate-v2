package main

import (
	"context"
	"fmt"
	"go-api-boilerplate/internal/dao"
	"go-api-boilerplate/internal/db"
	"os"

	"gorm.io/gorm"

	"go-api-boilerplate/pkg/common/logger"
)

var ctx = context.Background()
var seedDataFilePath = "internal/db/seed/seed.data.json"

func commonInit() {
	logger.Init()
	db.InitDB()
}

func main() {
	commonInit()
    logger.Logger.Info("Seeding successfully initialized")

    seedData, err := LoadSeedData(seedDataFilePath)
    if err != nil {
        logger.Logger.Error("Failed to load seed data:", "error", err)
        os.Exit(1)
    }
    logger.Logger.Info("Successfully loaded seed data")

    if err := seedDatabase(db.GetWriteDB(), seedData); err != nil {
        logger.Logger.Error("Failed to seed database:", "error", err)
        os.Exit(1)
    }

    logger.Logger.Info("Database seeding completed successfully")
}

func seedDatabase(db *gorm.DB, seedData *SeedData) error {
    tx := db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    seededUsers := make([]dao.User, len(seedData.Users))

    logger.Logger.Info("Seeding users...")
    for i, userData := range seedData.Users {
        user := dao.User{
            Name: userData.Name,
            Role: userData.Role,
        }
        
        if err := tx.Create(&user).Error; err != nil {
            tx.Rollback()
            return fmt.Errorf("failed to seed user %s: %w", user.Name, err)
        }
        seededUsers[i] = user
    }
    logger.Logger.Info("Successfully seeded users", "count", len(seededUsers))

    // Commit the transaction
    if err := tx.Commit().Error; err != nil {
        return fmt.Errorf("failed to commit transaction: %w", err)
    }

    return nil
}