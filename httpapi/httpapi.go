package httpapi

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"

)

func InitHttpRouter() *gin.Engine {
	router := gin.Default()


	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		MaxAge:           12 * time.Hour,
	}))
	router.POST("/device/power-status", handlePowerStatus)
	router.POST("/device/position", handleDevicePosition)
	router.POST("/device/status", handleDeviceStatus)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	return router
}

func handlePowerStatus(c *gin.Context) {
	var body struct {
		DeviceId string `json:"deviceId"`
		Status   string `json:"status"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Failed to parse power status request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received status from device %s: %s", body.DeviceId, body.Status)

	c.JSON(http.StatusOK, gin.H{
		"message": "status received",
		"device":  body.DeviceId,
		"status":  body.Status,
	})
}

func handleDevicePosition(c *gin.Context) {
	var body struct {
		DeviceId string  `json:"deviceId"`
		Lat      float64 `json:"lat"`
		Long     float64 `json:"long"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Failed to parse position request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received position from device %s: Lat %.2f, Long %.2f", body.DeviceId, body.Lat, body.Long)

	c.JSON(http.StatusOK, gin.H{
		"message": "status received",
		"device":  body.DeviceId,
		"lat":     body.Lat,
		"long":    body.Long,
	})
}

func handleDeviceStatus(c *gin.Context) {
	var body struct {
		DeviceId string `json:"deviceId"`
		Kwh      string `json:"kwh"`
		Volt     string `json:"volt"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Failed to parse device status request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received status from device %s: Volt %s, Kwh %s", body.DeviceId, body.Volt, body.Kwh)

	c.JSON(http.StatusOK, gin.H{
		"message": "status received",
		"device":  body.DeviceId,
		"kwh":     body.Kwh,
		"volt":    body.Volt,
	})
}
