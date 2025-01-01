package handler

import (
	"clinic/server/structures"
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

func (h *Handler) HandleMQTTMessage(client mqtt.Client, msg mqtt.Message) {
	fmt.Println("Received message: ", string(msg.Payload()))

	var data structures.IndicatorsStamp
	err := json.Unmarshal(msg.Payload(), &data)
	if err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return
	}

	data.Timestamp = time.Now().Format(time.RFC3339)
	err = h.services.IndicatorsStampAction.Create(data)
	if err != nil {
		log.Printf("Error saving data to DB: %v", err)
	}
}
