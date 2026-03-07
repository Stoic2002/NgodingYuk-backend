package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
	Port        string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		Port:        os.Getenv("PORT"),
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg
}

func InitDB(databaseURL string) *gorm.DB {
	dsn := databaseURL
	if len(dsn) > 13 && dsn[:13] == "postgresql://" {
		dsn = "postgres://" + dsn[13:]
	}

	// Disable pgx literal statement caching to prevent "prepared statement already exists" (42P05)
	// when sqlsandbox tests multiple dynamic schemas with the same query literal.
	if !strings.Contains(dsn, "statement_cache_capacity=0") {
		if strings.Contains(dsn, "?") {
			dsn += "&statement_cache_capacity=0"
		} else {
			dsn += "?statement_cache_capacity=0"
		}
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // Disable prepared statements for PgBouncer
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	sqlDB, err := db.DB()
	if err == nil {
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(5)
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(20)
	}

	log.Println("Database connected successfully")
	return db
}
