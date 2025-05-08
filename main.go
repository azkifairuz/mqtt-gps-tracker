package main

import (
	"github.com/azkifairuz/mqtt-gps-tracker/mqtt"
	"github.com/azkifairuz/mqtt-gps-tracker/httpapi"
)

func main() {
	go func() {
		mqtt.InitMqtt() 
	}()

	r := httpapi.InitHttpRouter()

	r.Run(":8080")
	

}