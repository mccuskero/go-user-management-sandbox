package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/mccuskero/go-user-management-sandbox/pkg/config"
	"github.com/mccuskero/go-user-management-sandbox/pkg/initializer"
	"github.com/mccuskero/go-user-management-sandbox/pkg/services"
)

var server *gin.Engine

func main() {
	// setup env configs
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	// connet to db
	initializer := initializer.NewInitializer()
	initializer.ConnectDB(&config)
	// TODO: maybe we don't need automigrate
	initializer.AutoMigrate()

	// Initialize the server
	server = gin.Default()

	// configure CORS, checking for cross site scripting...
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true
	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to User Management Sandbox"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	// setup services
	authService := services.NewAuthService()
	authService.Initialize(router, initializer)

	// startup server (report error if needed)
	log.Fatal(server.Run(":" + config.ServerPort))
}
