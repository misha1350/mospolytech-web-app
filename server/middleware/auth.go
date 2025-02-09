package middleware

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/misha1350/mospolytech-web-app/db/mysql"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(email string, password string) (map[string]interface{}, error) {
	ctx := context.Background()

	db, err := GetDB()
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
		"expires_at":   time.Now().Add(time.Hour * 24 * 14).Format("2006-01-02 15:04:05"),
	}

	return tokenDetails, nil
}

// This checks the tokenString sent over from the frontend and should return the user's details
func ValidateToken(c *gin.Context, tokenString string) (map[string]interface{}, error) {
	if tokenString == "" {
		return nil, errors.New("no token provided")
	}

	db, err := GetDB()
	if err != nil {
		return nil, fmt.Errorf("database error: %v", err)
	}
	queries := mysql.New(db)

	// Parse takes the token string and a function for looking up the key
	givenToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	if claims, ok := givenToken.Claims.(jwt.MapClaims); ok && givenToken.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, errors.New("token expired")
		}

		givenUserID := int32(claims["sub"].(float64))
		dbToken, err := queries.GetAuthToken(c, givenUserID)
		if err != nil {
			return nil, errors.New("user not found")
		}

		userDetails, err := queries.GetUserDataByID(c, dbToken.Userid)
		if err != nil {
			return nil, errors.New("user not found")
		}

		return map[string]interface{}{
			"email":     userDetails.Email,
			"firstname": userDetails.Firstname,
			"lastname":  userDetails.Lastname,
			"office":    userDetails.Officeid,
			"role":      userDetails.Roleid,
		}, nil
	}

	return nil, errors.New("invalid token")
}
