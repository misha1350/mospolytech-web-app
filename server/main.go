package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/misha1350/mospolytech-web-app/middleware"
)

func init() {
	// Initialize database connection
	if _, err := middleware.GetDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
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

	response, err := middleware.RegisterUser(userDetails)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(c.Writer, response, err.Error())
	} else {
		c.Writer.WriteHeader(http.StatusCreated)
		fmt.Fprint(c.Writer, response)
	}
}

func authenticationsHandler(c *gin.Context) {
	email, password := c.Request.FormValue("email"), c.Request.FormValue("password")
	if email != "" || password != "" {
		tokenDetails, err := middleware.GenerateToken(email, password)
		if err != nil {
			c.Writer.WriteHeader(http.StatusNotFound)
			fmt.Fprint(c.Writer, err.Error())
			return
		} else {
			c.SetSameSite(http.SameSiteLaxMode)
			c.SetCookie("Authorization", tokenDetails["token"].(string), 60*60*24*30, "/", "localhost", false, false)
			// Redirect the user back to the frontend
			c.Redirect(http.StatusMovedPermanently, "http://localhost:8087/client")
		}
	} else {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(c.Writer, "\nUsername and password are required to get a token.\n")
	}
}

func validationsHandler(c *gin.Context, token string) {
	userDetails, err := middleware.ValidateToken(c, token)
	// What userDetails looks like
	// userDetails := map[string]interface{}{
	// 	"email":     userDetails.Email,
	// 	"firstname": userDetails.Firstname,
	// 	"lastname":  userDetails.Lastname,
	// 	"office":    userDetails.Officeid,
	//  "role":		 userDetails.Roleid,
	// }
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"userDetails": userDetails})
	}
}

func main() {

	//logOutput()
	// A new "router" instance is created with this
	router := gin.New()
	router.Use(
		gin.Recovery(),
		middleware.Logger(),
	)

	router.StaticFile("/api/server/css/styles2.css", "./templates/css/styles2.css")

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*.html")

	// "router.GET" takes the URI to match the HTTP GET request, and the callback function in the form of a Gin Context struct to be executed.
	// "/api/server" is needed for Traefik to route the request to the server, according to the rule that you set in docker-compose.yml
	//
	router.GET("/api/server/ping", func(context *gin.Context) {
		// Check if the database connection is alive
		db, err := middleware.GetDB()
		if err != nil {
			context.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "error",
				"db":      "connection failed",
			})
			return
		}

		if err := db.Ping(); err != nil {
			context.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "error",
				"db":      "unreachable",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"db":      "connected",
		})
	})

	router.GET("/api/server/", func(context *gin.Context) {
		context.HTML(
			http.StatusOK,
			"index.html",
			gin.H{"title": "Home Page"})
	})

	router.GET("/api/server/login", func(context *gin.Context) {
		context.HTML(
			http.StatusOK,
			"login.html",
			gin.H{"title": "Login Page"})
		fmt.Println("ok")
	})

	router.GET("/api/server/img/logo_with_colors.png", func(context *gin.Context) {
		context.File("./templates/img/logo_with_colors.png")
	})

	router.GET("/api/server/img/logo_transparent.png", func(context *gin.Context) {
		context.File("./templates/img/logo_transparent.png")
	})

	router.POST("/api/server/login", func(context *gin.Context) {
		authenticationsHandler(context)
	})

	router.GET("/api/server/register", func(context *gin.Context) {
		context.HTML(
			http.StatusOK,
			"register.html",
			gin.H{"title": "Register Page"})
	})

	router.POST("/api/server/register", func(context *gin.Context) {
		registrationsHandler(context)
	})

	//TODO: Implement SELECT queries for the frontend
	router.GET("/api/server/get_users", middleware.GetUsers)

	//TODO: Complete user editing
	router.PUT("/api/server/user_edit", middleware.EditUser)

	router.POST("/api/server/check", func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		// // tokenString will contain the value of the Authorization header, which should be in the format "Bearer <token>"
		// // You can extract the token by splitting the string and taking the second part:
		// token := strings.Split(tokenString, " ")[1]
		validationsHandler(context, token)
	})

	// I had a great deal of trouble with this - apparently, when dockerizing the application and using Traefik to direct your requests,
	// you need to only specify the port number. not the entire "127.0.0.1:8086" address. Otherwise, Bad Gateway errors will be thrown.
	// This is due to Docker configuring the network for you, and the app receives connections from the likes of 172.18.0.1:8086.
	// Traefik knows better. You should too.
	srv := &http.Server{
		Addr:    ":8086",
		Handler: router,
	}

	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", srv.Addr)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context informs the server that it has 5 seconds to finish the currently handled request
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
