package config

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

var DB *gorm.DB
var Pool *pgxpool.Pool

func ConnectDB() {
    user := os.Getenv("PG_USER")
    host := os.Getenv("PG_HOST")
    port := os.Getenv("PG_PORT")
    password := os.Getenv("PG_PASSWORD")
    dbname := os.Getenv("PG_DATABASE")
	sslmode := os.Getenv("PG_SSLMODE")

    // Connection pool configuration
    maxConns := getEnvAsInt("PG_MAX_CONNS", 25)
    minConns := getEnvAsInt("PG_MIN_CONNS", 5)
    maxConnLifetime := getEnvAsDuration("PG_MAX_CONN_LIFETIME", "1h")
    maxConnIdleTime := getEnvAsDuration("PG_MAX_CONN_IDLE_TIME", "30m")

    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=10",
        host, port, user, password, dbname, sslmode,
    )

    log.Printf("Connecting to database at %s:%s with SSL mode=%s", host, port, sslmode)
    log.Printf("Pool config: maxConns=%d, minConns=%d, maxConnLifetime=%s, maxConnIdleTime=%s", 
               maxConns, minConns, maxConnLifetime, maxConnIdleTime)

    // Create pgxpool Config from DSN
    config, err := pgxpool.ParseConfig(dsn)
    if err != nil {
        log.Fatalf("Unable to parse DSN for pool: %v", err)
    }

    // Configure connection pool settings
    config.MaxConns = int32(maxConns)
    config.MinConns = int32(minConns)
    config.MaxConnLifetime = maxConnLifetime
    config.MaxConnIdleTime = maxConnIdleTime
    config.HealthCheckPeriod = 1 * time.Minute

    // ✅ CRITICAL: Disable prepared statements at the pgx level to prevent caching issues
    config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

    // ✅ Force IPv4 by overriding ConnConfig DialFunc
    dialer := &net.Dialer{
        Timeout:   10 * time.Second,
        KeepAlive: 10 * time.Second,
    }
    config.ConnConfig.DialFunc = func(ctx context.Context, network, addr string) (net.Conn, error) {
        return dialer.DialContext(ctx, "tcp4", addr) // force IPv4
    }

    // Create connection pool
    ctx := context.Background()
    pool, err := pgxpool.NewWithConfig(ctx, config)
    if err != nil {
        log.Fatalf("Unable to create connection pool: %v", err)
    }

    // Test the pool connection
    if err := pool.Ping(ctx); err != nil {
        log.Fatalf("Unable to ping database: %v", err)
    }

    Pool = pool

    // Create GORM DB instance using the pool
    sqlDB := stdlib.OpenDBFromPool(pool)

    db, err := gorm.Open(postgres.New(postgres.Config{
        Conn: sqlDB,
        // Disable prepared statements to fix "prepared statement already exists" error
        // This is necessary when using Supabase transaction pooling
        PreferSimpleProtocol: true,
    }), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
        // Additional GORM configuration to handle connection pooling better
        PrepareStmt:                              false, // Disable prepared statement caching
        DisableForeignKeyConstraintWhenMigrating: true,  // Additional safety for pooled connections
    })
    if err != nil {
        log.Fatalf("Failed to connect to database with GORM: %v", err)
    }

    DB = db
    log.Println("✅ Successfully connected to database with connection pooling over IPv4")
}

// Helper function to get environment variable as integer with default value
func getEnvAsInt(name string, defaultVal int) int {
    valueStr := os.Getenv(name)
    if value, err := strconv.Atoi(valueStr); err == nil {
        return value
    }
    return defaultVal
}

// Helper function to get environment variable as duration with default value
func getEnvAsDuration(name string, defaultVal string) time.Duration {
    valueStr := os.Getenv(name)
    if valueStr == "" {
        valueStr = defaultVal
    }
    if value, err := time.ParseDuration(valueStr); err == nil {
        return value
    }
    if defaultDuration, err := time.ParseDuration(defaultVal); err == nil {
        return defaultDuration
    }
    return 1 * time.Hour // fallback
}

// ClosePool closes the connection pool gracefully
func ClosePool() {
    if Pool != nil {
        Pool.Close()
        log.Println("🔌 Database connection pool closed")
    }
}

// GetPoolStats returns connection pool statistics for monitoring
func GetPoolStats() map[string]interface{} {
    if Pool == nil {
        return map[string]interface{}{
            "error": "Pool not initialized",
        }
    }
    
    stats := Pool.Stat()
    return map[string]interface{}{
        "max_conns":               stats.MaxConns(),
        "acquired_conns":          stats.AcquiredConns(),
        "idle_conns":             stats.IdleConns(),
        "constructing_conns":     stats.ConstructingConns(),
        "total_acquire_duration": stats.AcquireDuration().String(),
        "total_acquired_conns":   stats.AcquireCount(),
        "total_canceled_acquire": stats.CanceledAcquireCount(),
        "total_empty_acquire":    stats.EmptyAcquireCount(),
    }
}
