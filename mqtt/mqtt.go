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
	
	subscribeToTopic("device/device-power-status", getDevicePowerStatus)
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
	client.Publish("device/device-power-status", 0, false, respJSON)
}
