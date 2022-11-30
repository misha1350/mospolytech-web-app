package server

import "github.com/gin-gonic/gin"

func Server() {
	server := gin.Default()
	server.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run(":8080") // listen and serve on 0.0.0.0:8080
}
