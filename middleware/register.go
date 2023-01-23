package middleware

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func RegisterUser(userDetails map[string]interface{}) (string, error) {
	// ctx := context.Background()

	// db, err := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/session1_xx")
	// if err != nil {
	// 	return "", err
	// }

	// queries := mysql.New(db)

	office := userDetails["officeid"].(int32)
	log.Printf("%v, %T", office, office)
	// result, err := queries.AddUser(ctx, mysql.AddUserParams{
	// 	Email:     userDetails["email"].(string),
	// 	Password:  userDetails["password"].(string),
	// 	Firstname: userDetails["firstname"].(string),
	// 	Lastname:  userDetails["lastname"].(string),
	// 	Officeid:  userDetails["officeid"].(int32),
	// 	Birthdate: userDetails["birthdate"].(time.Time),
	// })
	// Use addUser function to add user to database from db/mysql/query.sql
	// queryString := "insert into system_users(username, password) values (?, ?)"
	// stmt, err := db.Prepare(queryString)
	// if err != nil {
	// 	return "", err
	// }
	// defer stmt.Close()

	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	// _, err = stmt.Exec(username, hashedPassword)
	// if err != nil {
	// 	return "", err
	// }
	return "Success\r\n", nil

}

// func RegisterUser2() error {
// 	db, err := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/session1_xx")
// 	if err != nil {
// 		return err
// 	}
// }
