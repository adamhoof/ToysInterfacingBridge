package main

import (
	"github.com/adamhoof/ToysInterfacingBridge/pkg/Buttons"
	"github.com/adamhoof/ToysInterfacingBridge/pkg/Database"
	"github.com/adamhoof/ToysInterfacingBridge/pkg/Env"
	"github.com/adamhoof/ToysInterfacingBridge/pkg/Keyboards"
	"github.com/adamhoof/ToysInterfacingBridge/pkg/MQTTs"
	"github.com/adamhoof/ToysInterfacingBridge/pkg/TelegramBot"
	"github.com/adamhoof/ToysInterfacingBridge/pkg/Toy"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
	"sync"
)

var me = TelegramBot.MeUser{}

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
	me.Id = os.Getenv("telegramID")

	db := Database.PostgresDatabase{}
	go func() {
		db.Connect()
		db.TestConnection()
	}()

	toyCommandBot := TelegramBot.ToyCommandBot{}
	go func() {
		toyCommandBot.SetToken("Auth/BotToken")
		toyCommandBot.StartBot()
	}()

	mqttClient := MQTTs.CreateClient(MQTTs.GetClientConfig())

	toyBag := make(map[string]Toy.Toy)
	db.PullToysData(toyBag)

	buttonFactory := Buttons.Factory{ToyButtonTemplates: map[string]string{
		"on":     "⬜",
		"white":  "⬜",
		"yellow": "\U0001F7E8",
		"blue":   "\U0001F7E6",
		"green":  "\U0001F7E9",
		"red":    "\U0001F7E5",
		"pink":   "\U0001F7EA",
		"orange": "\U0001F7E7",
		"off":    "🚫",
		"1":      "🌞",
		"0":      "🌚"}}
	keyboardFactory := Keyboards.KeyboardFactory{MenuButtonTemplates: make(map[string]string)}

	for _, toy := range toyBag {
		buttonFactory.GenerateToyCommandButtons(toy.Buttons, toy.ID, toy.AvailableCommands)
		toy.Keyboard = keyboardFactory.GenerateToyCommandsKeyboard(toy.Buttons)
	}

	MQTTs.ConnectClient(&mqttClient)

	routineSyncer := sync.WaitGroup{}
	routineSyncer.Add(1)
	routineSyncer.Wait()
}
