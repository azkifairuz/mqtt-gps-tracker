package main

import (
	"time"
	"github.com/azkifairuz/mqtt-gps-tracker/mqtt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	go func() {
		mqtt.InitMqtt() 
	}()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		MaxAge:           12 * time.Hour,
	}))

	// Test route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.Run(":8080")
}