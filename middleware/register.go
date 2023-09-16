package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/misha1350/mospolytech-web-app/db/mysql"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(userDetails map[string]interface{}) (string, error) {
	ctx := context.Background()

	db, err := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/session1_xx")
	if err != nil {
		return "", err
	}
	queries := mysql.New(db)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(fmt.Sprint(userDetails["password"])), 14)

	parsedOfficeID, err := strconv.ParseInt(userDetails["office"].(string), 10, 32)
	if err != nil {
		return "", err
	}

	parsedBirthdate, err := time.Parse("2006-01-02", userDetails["birthdate"].(string)) // adjust the layout string according to the format of your date string
	if err != nil {
		return "", err
	}

	queries.AddUser(ctx, mysql.AddUserParams{
		// RoleID is set to 1 by default in "db/query.sql"
		Email:     userDetails["email"].(string),
		Password:  string(hashedPassword),
		Firstname: userDetails["firstname"].(string),
		Lastname:  userDetails["lastname"].(string),
		Officeid:  int32(parsedOfficeID),
		Birthdate: parsedBirthdate,
		// Active is set to 1 by default in "db/query.sql"
	})
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}
	return "Success\r\n", nil
}
