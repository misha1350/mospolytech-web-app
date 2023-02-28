package main

import (
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
	// // persistent logs can grow extremely large with this:
	// f, _ := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func registrationsHandler(c *gin.Context) {
	userDetails := map[string]interface{}{
		"email":     c.Request.FormValue("email"),
		"firstname": c.Request.FormValue("firstname"),
		"lastname":  c.Request.FormValue("lastname"),
		"birthdate": c.Request.FormValue("birthdate"),
		"office":    c.Request.FormValue("office"),
		"password":  c.Request.FormValue("password"),
	}
	log.Println(userDetails["email"], userDetails["firstname"], userDetails["lastname"], userDetails["birthdate"], userDetails["office"], userDetails["password"])

	response, err := middleware.RegisterUser(userDetails)
	if err != nil {
		fmt.Fprint(c.Writer, err.Error())
	} else {
		fmt.Fprint(c.Writer, response)
	}
}

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

	//logOutput()
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

	server.GET("/img/DS2017_TP09_2_colors_with_bg_4x.png", func(context *gin.Context) {
		context.File("./templates/img/DS2017_TP09_2_colors_with_bg_4x.png")
	})

	server.POST("/login", func(context *gin.Context) {
		email := context.PostForm("email")
		password := context.PostForm("password")
		log.Println(context.Writer, "email is ", email, "and password is ", password)
		authenticationsHandler(context)
	})

	server.GET("/register", func(context *gin.Context) {
		context.HTML(
			http.StatusOK,
			"register.html",
			gin.H{"title": "Register Page"})
	})

	server.POST("/register", func(context *gin.Context) {
		registrationsHandler(context)
	})

	server.Run("127.0.0.1:8082") // listen and serve on this address
}
