package MQTTs

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
)

func CreateClient(options mqtt.ClientOptions) mqtt.Client {
	return mqtt.NewClient(&options)
}

func ConnectClient(client *mqtt.Client) {
	if token := (*client).Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("failed to connect mqtt client: %v", token.Error())
	}
	fmt.Println("Connected to MQTTs server")
}

func SetSubscription(client *mqtt.Client, topic string, incomingMessageHandler mqtt.MessageHandler) {
	if token := (*client).Subscribe(topic, 0, incomingMessageHandler); token.Wait() && token.Error() != nil {
		log.Fatalf("failed to create subscription: %v", token.Error())
	}
}

func Publish(client *mqtt.Client, topic string, payload string) {
	if token := (*client).Publish(topic, 0, true, payload); token.Wait() && token.Error() != nil {
		log.Fatalf("failed to publish text: %v", token.Error())
	}
}
