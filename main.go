package main

import (

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"

	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/misha1350/mospolytech-web-app/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func logOutput() {
	f, _ := os.Create("access.log")
	// persistent logs can grow extremely large with this:
	// f, _ := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

// func registrationsHandler(c *gin.Context) {
// 	c.Request.ParseForm()
// 	if c.Request.FormValue("email") == "" || c.Request.FormValue("password") == "" {
// 		fmt.Fprintf(c.Writer, "Please enter a valid email and password.\r\n")
// 	} else {
// 		response, err := middleware.RegisterUser(c.Request.FormValue("email"), c.Request.FormValue("password"))
// 		if err != nil {
// 			fmt.Fprint(c.Writer, err.Error())
// 		} else {
// 			fmt.Fprint(c.Writer, response)
// 		}
// 	}
// }

func authenticationsHandler(c *gin.Context) {
	email, password := c.Request.FormValue("email"), c.Request.FormValue("password")
	if email != "" || password != "" {
		tokenDetails, err := middleware.GenerateToken(email, password)
		if err != nil {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(c.Writer, err.Error())
			return
		} else {
			enc := json.NewEncoder(c.Writer)
			enc.SetIndent("", "  ")
			enc.Encode(tokenDetails)
		}
	} else {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(c.Writer, "\nUsername and password are required to get a token.\n")
	}
}

// func testResourceHandler(c *gin.Context) {
// 	authToken := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")[1]
// 	userDetails, err := middleware.ValidateToken(authToken)
// 	if err != nil {
// 		fmt.Fprintf(c.Writer, err.Error())
// 	} else {
// 		// E-mail or First and Last Name?
// 		username := fmt.Sprint(userDetails["email"])
// 		fmt.Fprintf(c.Writer, "Welcome, "+username+"\r\n")
// 	}
// }

func main() {

	logOutput()
	server := gin.New()
	server.Use(
		gin.Recovery(),
		middleware.Logger(),
		// middleware.BasicAuth(),
	)

	//serve css
	server.StaticFile("/css/", "./templates/css/index.css")

	//serve javascript
	server.StaticFile("/js/script.js", "./templates/js/script.js")

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	server.LoadHTMLGlob("templates/*.html")

	// gindump.Dump()

	server.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/", func(context *gin.Context) {
		context.HTML(
			http.StatusOK,
			"index.html",
			gin.H{"title": "Home Page"})
	})

	server.GET("/login", func(context *gin.Context) {
		context.HTML(
			http.StatusOK,
			"login.html",
			gin.H{"title": "Login Page"})
		fmt.Println("ok")
	})

	server.GET("/register", func(context *gin.Context) {
		context.HTML(
			http.StatusOK,
			"register.html",
			gin.H{"title": "Register Page"})
	})

	server.POST("/login", func(context *gin.Context) {
		email := context.PostForm("email")
		password := context.PostForm("password")
		log.Println(context.Writer, "email is %s and password is %s", email, password)
		authenticationsHandler(context)
	})
	// useless code for now
	// server.POST("/login", func(context *gin.Context) {
	// 	login := context.PostForm("login")
	// 	password := context.PostForm("password")
	// 	fmt.Fprintf(context.Writer, "login is %s and password is %s", login, password)
	// 	fmt.Fprintf(context.Writer, "\nshould return '%T' and '%T'", login, password)
	// })

	server.Run("127.0.0.1:8082") // listen and serve on this address

	// sqlc connection to database method

	// GORM connection to database method

	// dsn := "root:123qwe@tcp(127.0.0.1:3306/session1_xx)"
	// gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("Unable to open connection to the database.")
	// }

	// database/sql connection to database method

	// fmt.Println("Using SQL driver:", sql.Drivers())
	// db, err := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/session1_xx")
	// if err != nil {
	// 	log.Fatal("Unable to open connection to the database.")
	// }
	// defer db.Close()
}
