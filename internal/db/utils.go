package db

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func Health(db *gorm.DB) map[string]string {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()

    stats := make(map[string]string)

    sqlDB, err := db.DB()
    if err != nil {
        stats["status"] = "down"
        stats["error"] = fmt.Sprintf("failed to get sql.DB instance: %v", err)
        log.Fatalf("failed to get sql.DB instance: %v", err)
        return stats
    }

    // Ping the database
    err = sqlDB.PingContext(ctx)
    if err != nil {
        stats["status"] = "down"
        stats["error"] = fmt.Sprintf("db down: %v", err)
        log.Fatalf("db down: %v", err)
        return stats
    }

    // Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "Database connection healthy"

	// Get database stats
    dbStats := sqlDB.Stats()
    stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
    stats["in_use_connections"] = strconv.Itoa(dbStats.InUse)
    stats["idle_connections"] = strconv.Itoa(dbStats.Idle)
    stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
    stats["wait_duration_ms"] = strconv.FormatInt(dbStats.WaitDuration.Milliseconds(), 10)
    stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
    stats["max_idle_time_closed"] = strconv.FormatInt(dbStats.MaxIdleTimeClosed, 10)
    stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

    // Get connection pool settings for evaluation
    maxOpenConns := dbStats.MaxOpenConnections
    if maxOpenConns <= 0 {
        maxOpenConns = 25 // Default Go sql.DB max open connections
    }

    // Evaluate stats to provide a health message
    utilizationPercent := float64(dbStats.InUse) / float64(maxOpenConns) * 100
    if utilizationPercent > 80 {
        stats["message"] = fmt.Sprintf("The database is experiencing heavy load (%.1f%% utilization).", utilizationPercent)
    }

    if dbStats.WaitCount > 1000 {
        stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
    }

    if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
        stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
    }

    if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
        stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
    }     

    return stats
}