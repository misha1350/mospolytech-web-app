package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func init() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func DbConnect() (*sql.DB, error) {
	return DB, nil
}

func DbPing() error {
	err := DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

// func main() {
// 	db, err := DbConnect()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Use the db connection
// 	// ...

// 	db.Close()
// }
