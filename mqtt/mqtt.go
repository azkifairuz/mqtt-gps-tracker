package mqtt

import (
	"encoding/json"
	"log"

	// "time"

	// "github.com/azkifairuz/rfid-presence-api/controllers"
	// "github.com/azkifairuz/rfid-presence-api/helper"
	// "github.com/azkifairuz/rfid-presence-api/initializers"
	"github.com/azkifairuz/rfid-presence-api/helper"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)
var mqttClient mqtt.Client

func InitMqtt(){
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")

	mqttClient = mqtt.NewClient(opts)
	token := mqttClient.Connect()
	token.Wait()
	if token.Error() != nil {
		log.Fatalf("failed to connect to mqtt broker: %v",token.Error())
	}
	log.Printf("connected to mqtt broker")
	
	subscribeToTopic("device/power-status", getDevicePowerStatus)
	subscribeToTopic("device/position", getDeviceLocation)
	subscribeToTopic("device/status", getDeviceStatus)
	// subscribeToTopic("rfid-system/read", ReadCardFromMQTT)
}

func subscribeToTopic(topic string, handler mqtt.MessageHandler)  {
	token := mqttClient.Subscribe(topic,1 ,handler)
	token.Wait()
	if token.Error() != nil {
		log.Fatalf("failed to connect to mqtt broker: %v",token.Error())
	}
	log.Printf("connected to mqtt broker")

}

func getDevicePowerStatus(client mqtt.Client, msg mqtt.Message) {
	var body struct {
		DeviceId string `json:"deviceId"`
		Status     string `json:"status"`
	}

	if err := helper.ParseJSON(msg.Payload(), &body); err != nil {
		log.Println("Failed to parse register card message:", err)
		return
	}

	log.Printf("Received status from device %s: %s", body.DeviceId, body.Status)


	response := map[string]string{
		"message": "status received",
		"device":  body.DeviceId,
		"status": body.Status,
	}
	respJSON, _ := json.Marshal(response)
	client.Publish("device/power-status/ack", 0, false, respJSON)
}

func getDeviceLocation(client mqtt.Client, msg mqtt.Message) {
	var body struct {
		DeviceId string `json:"deviceId"`
	 	Lat      float64 `json:"lat"`   // contoh: -6.200000
		Long     float64 `json:"long"`
	}

	if err := helper.ParseJSON(msg.Payload(), &body); err != nil {
		log.Println("Failed to parse register card message:", err)
		return
	}

	log.Printf("Received status from device %s: %.2f", body.DeviceId, body.Long)

	response := map[string]interface{}{
		"message": "status received",
		"device":  body.DeviceId,
		"lat":     body.Lat,
		"long":    body.Long,
	}
	respJSON, _ := json.Marshal(response)
	client.Publish("device/position/ack", 0, false, respJSON)
}

func getDeviceStatus(client mqtt.Client, msg mqtt.Message) {
	var body struct {
		DeviceId string `json:"deviceId"`
	 	Kwh      string `json:"kwh"`  
		Volt     string `json:"volt"`
	}

	if err := helper.ParseJSON(msg.Payload(), &body); err != nil {
		log.Println("Failed to parse register card message:", err)
		return
	}

	log.Printf("Received status from device %s: %s: %s", body.DeviceId, body.Volt,body.Kwh)

	response := map[string]string{
		"message": "status received",
		"device":  body.DeviceId,
		"kwh":     body.Kwh,
		"volt":    body.Volt,
	}
	respJSON, _ := json.Marshal(response)
	client.Publish("device/status/ack", 0, false, respJSON)
}