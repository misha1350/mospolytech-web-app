package middleware

import (
	"context"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/misha1350/mospolytech-web-app/server/db/mysql"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(userDetails map[string]interface{}) (string, error) {
	ctx := context.Background()

	db, err := DbConnect()
	if err != nil {
		return "Database connection failed", err
	}
	queries := mysql.New(db)

	_, err = queries.GetUserData(ctx, userDetails["email"].(string))
	if err == nil {
		return "User already exists", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprint(userDetails["password"])), 14)
	if err != nil {
		return "Password hashing failed", err
	}
	parsedOfficeID, err := strconv.ParseInt(userDetails["office"].(string), 10, 32)
	if err != nil {
		return "OfficeID parsing failed", err
	}

	parsedBirthdate, err := time.Parse("2006-01-02", userDetails["birthdate"].(string)) // adjust the layout string according to the format of your date string
	if err != nil {
		return "Birthdate parsing failed", err
	}

	_, err = queries.AddUser(ctx, mysql.AddUserParams{
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
		return "User creation failed, looks like the developer will have a sleepless night", err
	}

	return "Success\r\n", nil
}
