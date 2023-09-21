package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DbConnect()
}

func DbConnect() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	DB, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	return DB, nil
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
