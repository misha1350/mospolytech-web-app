package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/misha1350/mospolytech-web-app/db/mysql"
)

func EditUser(c *gin.Context) {
	type editedUserInfo struct {
		Email     string `json:"email"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Office    string `json:"office"`
		Role      string `json:"role"`
	}

	var editedUser editedUserInfo
	if err := c.ShouldBindJSON(&editedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := GetDB()
	if err != nil {
		c.JSON(500, gin.H{"error": "Database connection failed"})
		return
	}
	queries := mysql.New(db)

	fmt.Println(editedUser.Email)
	fmt.Println(editedUser.FirstName)
	fmt.Println(editedUser.LastName)
	fmt.Println(editedUser.Office)
	fmt.Println(editedUser.Role)

	// TODO: Get User ID from the user SELECTed to be edited from the frontend
	err = queries.UpdateUser(c, mysql.UpdateUserParams{
		Roleid:    1,
		Email:     editedUser.Email,
		Firstname: editedUser.FirstName,
		Lastname:  editedUser.LastName,
		Officeid:  1,
		ID:        1,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
}
