package main

import (
	"github.com/adamhoof/ToysInterfacingBridge/pkg/Database"
	"github.com/adamhoof/ToysInterfacingBridge/pkg/Env"
	"github.com/adamhoof/ToysInterfacingBridge/pkg/MQTTs"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"sync"
)

func defaultResponseHandler(toyName string, db Database.Database) func(client mqtt.Client, message mqtt.Message) {
	handler := func(client mqtt.Client, message mqtt.Message) {

		func() {
			/*receivedMessage := string(message.Payload())
			services.db.UpdateToyMode(toyName, receivedMessage)
			_, err := services.botHandler.bot.Send(&me, toyName+": "+receivedMessage)
			if err != nil {
				return
			}*/

		}()
	}
	return handler
}

func main() {
	Env.SetVariables()

	db := Database.PostgresDatabase{}

	var routineSyncer sync.WaitGroup
	routineSyncer.Add(1)
	go func(routineSyncer *sync.WaitGroup) {
		defer routineSyncer.Done()
		db.Connect()
		db.TestConnection()
	}(&routineSyncer)

	routineSyncer.Add(1)
	go func(routineSyncer *sync.WaitGroup) {
		defer routineSyncer.Done()
		db.Connect()
		db.TestConnection()
	}(&routineSyncer)

	mqttClient := MQTTs.CreateClient(MQTTs.GetClientConfig())
	MQTTs.ConnectClient(&mqttClient)

	routineSyncer.Wait()
}
