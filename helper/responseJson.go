package helper

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

// Helper function for sending JSON responses
func ResponseDefault(c *gin.Context, statusCode int, data interface{}, message string) {
	c.JSON(statusCode, gin.H{
		"data":    data,
		"message": message,
		"statusCode": statusCode,
	})
}

func ParseJSON(data []byte, target interface{}) error {
	if len(data) == 0 {
		return errors.New("payload is empty")
	}
	if err := json.Unmarshal(data, target); err != nil {
		return err
	}
	return nil
}