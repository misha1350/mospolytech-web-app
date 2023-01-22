package middleware

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username string, password string) (string, error) {
	db, err := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/session1_xx")
	if err != nil {
		return "", err
	}
	// Use addUser function to add user to database from db/mysql/query.sql
	queryString := "insert into system_users(username, password) values (?, ?)"
	stmt, err := db.Prepare(queryString)
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	_, err = stmt.Exec(username, hashedPassword)
	if err != nil {
		return "", err
	}
	return "Success\r\n", nil

}

// func RegisterUser2()
