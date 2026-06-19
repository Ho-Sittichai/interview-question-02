package main

import (
	"log"
	"path/filepath"
	"runtime"

	"auth-login/database"
	"auth-login/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Resolve path to shared database relative to this file's location
	_, filename, _, _ := runtime.Caller(0)
	dbPath := filepath.Join(filepath.Dir(filename), "../shared/app.db")

	database.Init(dbPath)

	r := gin.Default()

	// CORS for local development
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api/auth")
	{
		// This backend handles LOGIN ONLY
		api.POST("/login", handlers.Login)
	}

	log.Println("Go Login Backend running on :5002")
	if err := r.Run(":5002"); err != nil {
		log.Fatal(err)
	}
}
