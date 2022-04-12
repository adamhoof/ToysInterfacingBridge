package main

import (
	"github.com/adamhoof/ToysInterfacingBridge/pkg/Database"
	"github.com/adamhoof/ToysInterfacingBridge/pkg/Env"
	"github.com/adamhoof/ToysInterfacingBridge/pkg/MQTTs"
	"github.com/adamhoof/ToysInterfacingBridge/pkg/TelegramBot"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
)

var me = TelegramBot.MeUser{Id: os.Getenv("adminUser")}

func defaultResponseHandler(toyName string, db Database.Database, bot TelegramBot.TelegramBot) func(client mqtt.Client, message mqtt.Message) {
	handler := func(client mqtt.Client, message mqtt.Message) {
		func() {
			receivedMessage := string(message.Payload())
			db.UpdateToyMode(toyName, receivedMessage)
			bot.SendTextMessage(&me, toyName+": "+receivedMessage)
		}()
	}
	return handler
}

func main() {
	Env.SetVariables()

	db := Database.PostgresDatabase{}
	go func() {
		db.Connect()
		db.TestConnection()
	}()

	mqttClient := MQTTs.CreateClient(MQTTs.GetClientConfig())
	MQTTs.ConnectClient(&mqttClient)

	cmdBot := TelegramBot.ToyCommandBot{}
	cmdBot.SetToken("Auth/BotToken")
	cmdBot.StartBot()
}
