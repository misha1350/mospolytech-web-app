package main

import (

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"

	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server := gin.Default()
	//serve css
	server.StaticFile("/css/", "./templates/css/index.css")

	//serve javascript
	server.StaticFile("/js/script.js", "./templates/js/script.js")

	// i forgor if it reads and does anything with tmpl's 💀💀💀
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery()) // middlewares.Logger(),
	// middlewares.BasicAuth(),
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

	// server.GET("/js/script.js", func(context *gin.Context) {
	// 	if pusher := context.Writer.Pusher(); pusher != nil {
	// 		// use pusher.Push() to do server push
	// 		if err := pusher.Push("/js/script.js", nil); err != nil {
	// 			log.Printf("Failed to push: %v", err)
	// 		}
	// 	}
	// })

	server.Run("127.0.0.1:8082") // listen and serve on 0.0.0.0:8080

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
