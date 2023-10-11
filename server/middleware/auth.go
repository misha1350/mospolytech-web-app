package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/misha1350/mospolytech-web-app/db/mysql"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(email string, password string) (map[string]interface{}, error) {
	ctx := context.Background()

	db, err := DbConnect()
	if err != nil {
		return nil, err
	}
	queries := mysql.New(db)

	// lookup returns struct with Email and Password (strings)
	user, err := queries.GetUserData(ctx, email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid username or password.\r")
	} else {
		fmt.Println("user is authenticated.\r")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, errors.New("failed to generate token.\r")
	}

	err = queries.InsertAuthToken(ctx, mysql.InsertAuthTokenParams{
		Userid:      user.ID,
		Authtoken:   tokenString,
		Generatedat: time.Now(),
		Expiresat:   time.Now().Add(time.Hour * 24 * 7),
	})

	if err != nil {
		return nil, errors.New("failed to write token to database.\r")
	}

	tokenDetails := map[string]interface{}{
		"token_type":   "Bearer",
		"token":        tokenString,
		"generated_at": time.Now().Format("2006-01-02 15:04:05"),
		"expires_at":   time.Now().Add(time.Hour * 24 * 30).Format("2006-01-02 15:04:05"),
	}

	return tokenDetails, nil
}

// This checks the tokenString sent over from the frontend and should return the user's details
func ValidateToken(c *gin.Context, tokenString string) (map[string]interface{}, error) {
	//TODO: Validate tokens using the database and a cookie
	db, err := DbConnect()
	if err != nil {
		return nil, err
	}
	queries := mysql.New(db)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return nil, &os.LinkError{}
	}
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	givenToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "cannot parse token"})
		// c.Abort()
		return nil, err
	}

	fmt.Println("Provided token's UserID is:", givenToken.Claims.(jwt.MapClaims)["sub"])

	if claims, ok := givenToken.Claims.(jwt.MapClaims); ok && givenToken.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, errors.New("token expired")
		}

		givenUserID := int32(claims["sub"].(float64))

		// Find the dbToken.UserID in database from provided token
		dbToken, err := queries.GetAuthToken(c, givenUserID)
		if err != nil {
			return nil, errors.New("user not found")
		}

		fmt.Println("dbToken:", dbToken)
		// fmt.Println("user:", tokenString)
		fmt.Println("givenToken:", givenToken)

		userDetails, err := queries.GetUserDataByID(c, dbToken.Userid)
		if err != nil {
			return nil, errors.New("user not found")
		}

		userDetailsMap := map[string]interface{}{
			"email":     userDetails.Email,
			"firstname": userDetails.Firstname,
			"lastname":  userDetails.Lastname,
			"office":    userDetails.Officeid,
			"password":  userDetails.Password,
		}

		return userDetailsMap, nil
	} else {
		return nil, errors.New("unauthorized")
	}
}

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
// err = stmt.QueryRow(tokenString).Scan(&userId, &username, &generatedAt, &expiresAt)
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
