package main

import (

	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server := gin.Default()
	server.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run(":8080") // listen and serve on 0.0.0.0:8080

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
