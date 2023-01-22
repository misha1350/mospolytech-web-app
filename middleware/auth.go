package middleware

import (
	"context"
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/misha1350/mospolytech-web-app/db/mysql"
)

// func BasicAuth() gin.HandlerFunc {
// 	return gin.BasicAuth(gin.Accounts{
// 		"admin": "admin123",
// 	})
// }

func GenerateToken(email string, password string) (map[string]interface{}, error) {
	ctx := context.Background()

	db, err := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/session1_xx")
	if err != nil {
		return nil, err
	}

	queries := mysql.New(db)

	// returns struct with ID (int32) and Password (string)
	user, err := queries.GetUsernameAndPassword(ctx, email)
	if err != nil {
		return nil, err
	}
	log.Println(user)
	log.Println(user.Email, user.Password)
	// Does user exist in the database?
	// queryString := "select user_id, password from system_users where username = ?"
	// stmt, err := db.Prepare(queryString)
	// if err != nil {
	// 	return nil, err
	// }

	// defer stmt.Close()
	// userId := 0
	// accountPassword := ""
	// err = stmt.QueryRow(username).Scan(&userId, &accountPassword)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return nil, errors.New("Invalid username or password.\r\n")
	// 	}
	// 	return nil, err
	// }
	// compare md5 hashes
	hash := md5.Sum([]byte(password))
	fmt.Printf("Supplied password's hash:\t%v\n", string(hash[:]))
	fmt.Printf("User's E-Mail:\t\t\t%v\n", user.Email)
	fmt.Printf("User's password's hash:\t\t%v\n", user.Password)
	if user.Password != string(hash[:]) {
		fmt.Println("Invalid username or password.\r")
		// return nil, errors.New("Invalid username or password.\r\n")
	} else {
		fmt.Println("User is authenticated.\r")
	}
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(hash))
	// if err != nil {
	// 	return nil, errors.New("Invalid username or password.\r\n")
	// }
	// Generate token
	// queryString = "insert into authentication_tokens(user_id, auth_token, generated_at, expires_at) values (?, ?, ?, ?)"
	// stmt, err = db.Prepare(queryString)
	// if err != nil {
	// 	return nil, err
	// }
	// defer stmt.Close()
	// randomToken := make([]byte, 32)
	// _, err = rand.Read(randomToken)
	// if err != nil {
	// 	return nil, err
	// }
	// authToken := base64.URLEncoding.EncodeToString(randomToken)
	// const timeLayout = "2006-01-02 15:04:05"
	// dt := time.Now()
	// expirtyTime := time.Now().Add(time.Minute * 60)
	// generatedAt := dt.Format(timeLayout)
	// expiresAt := expirtyTime.Format(timeLayout)
	// _, err = stmt.Exec(userId, authToken, generatedAt, expiresAt)
	// if err != nil {
	// 	return nil, err
	// }
	// tokenDetails := map[string]interface{}{
	// 	"token_type":   "Bearer",
	// 	"auth_token":   authToken,
	// 	"generated_at": generatedAt,
	// 	"expires_at":   expiresAt,
	// }
	return nil, nil
	//}
	// func ValidateToken(authToken string) (map[string]interface{}, error) {
	// 	db, err := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/session1_xx")
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// queryString := `select
	//             system_users.user_id,
	//             username,
	//             generated_at,
	//             expires_at
	//         from authentication_tokens
	//         left join system_users
	//         on authentication_tokens.user_id = system_users.user_id
	//         where auth_token = ?`
	// stmt, err := db.Prepare(queryString)
	// if err != nil {
	// 	return nil, err
	// }
	// defer stmt.Close()
	// userId := 0
	// username := ""
	// generatedAt := ""
	// expiresAt := ""
	// err = stmt.QueryRow(authToken).Scan(&userId, &username, &generatedAt, &expiresAt)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return nil, errors.New("Invalid access token.\r\n")
	// 	}
	// 	return nil, err
	// }
	// const timeLayout = "2006-01-02 15:04:05"
	// expiryTime, _ := time.Parse(timeLayout, expiresAt)
	// currentTime, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))
	// if expiryTime.Before(currentTime) {
	// 	return nil, errors.New("The token is expired.\r\n")
	// }
	// userDetails := map[string]interface{}{
	// 	"user_id":      userId,
	// 	"username":     username,
	// 	"generated_at": generatedAt,
	// 	"expires_at":   expiresAt,
	// }
	//return userDetails, nil
}
