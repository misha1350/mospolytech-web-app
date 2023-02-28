package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

	office := userDetails["office"]
	log.Printf("%v, %T", office, office)

	//TODO:
	//fix the userDetails["office"] type (it's a string, but it should be int32)
	//fix the userDetails["birthdate"] type (it's a string, but (i guess) it should be time.Time)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(fmt.Sprint(userDetails["password"])), 14)
	fmt.Printf("debug: %v, %T", string(hashedPassword), string(hashedPassword))

	result, err := queries.AddUser(ctx, mysql.AddUserParams{
		Email:     userDetails["email"].(string),
		Password:  string(hashedPassword),
		Firstname: userDetails["firstname"].(string),
		Lastname:  userDetails["lastname"].(string),
		Officeid:  userDetails["office"].(sql.NullInt32),
		Birthdate: userDetails["birthdate"].(sql.NullTime),
	})
	if err != nil {
		return "", err
	}
	fmt.Println("debug RESULT: ", result)
	// _, err = stmt.Exec(username, hashedPassword)
	if err != nil {
		return "", err
	}
	return "Success\r\n", nil

}

// func RegisterUser2() error {
// 	db, err := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/session1_xx")
// 	if err != nil {
// 		return err
// 	}
// }
