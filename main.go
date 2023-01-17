package main

import (

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"

	"io"
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

func main() {

	logOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger())

	//serve css
	server.StaticFile("/css/", "./templates/css/index.css")

	//serve javascript
	server.StaticFile("/js/script.js", "./templates/js/script.js")

	// load HTML templates
	server.LoadHTMLGlob("templates/*.html")

	// middleware.BasicAuth(),
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

	server.Run("127.0.0.1:8082") // listen and serve on this address

	// sqlc

	// GORM

	// dsn := "root:123qwe@tcp(127.0.0.1:3306/session1_xx)"
	// gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("Unable to open connection to the database.")
	// }

	// database/sql

	// fmt.Println("Using SQL driver:", sql.Drivers())
	// db, err := sql.Open("mysql", "root:123qwe@tcp(127.0.0.1:3306)/session1_xx")
	// if err != nil {
	// 	log.Fatal("Unable to open connection to the database.")
	// }
	// defer db.Close()
}
