package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/misha1350/mospolytech-web-app/config"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

func initDB() error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Load database configuration
	dbConfig := config.LoadDatabaseConfig()

	// Set connection pool parameters from config
	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)
	db.SetConnMaxLifetime(dbConfig.ConnMaxLifetime)

	// Verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	return nil
}

// GetDB returns a database instance, initializing it if necessary
func GetDB() (*sql.DB, error) {
	var initErr error
	dbOnce.Do(func() {
		initErr = initDB()
	})
	if initErr != nil {
		return nil, initErr
	}
	return db, nil
}

// CloseDB closes the database connection
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
