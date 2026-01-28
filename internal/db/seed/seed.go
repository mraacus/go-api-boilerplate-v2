package main

import (
	"fmt"
	"os"

	"github.com/bytedance/sonic"
)

type SeedData struct {
	Users    []UserData   `json:"users"`
	Metadata Metadata     `json:"metadata"`
}

type UserData struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type Metadata struct {
	Version     string `json:"version"`
	Description string `json:"description"`
}

// Loads seed data from seed data JSON file
func LoadSeedData(filePath string) (*SeedData, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read seed data file: %w", err)
	}

	var seedData SeedData
	if err := sonic.Unmarshal(data, &seedData); err != nil {
		return nil, fmt.Errorf("failed to parse seed data JSON: %w", err)
	}

	return &seedData, nil
}
