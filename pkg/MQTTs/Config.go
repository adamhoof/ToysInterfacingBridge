package MQTTs

import (
	"github.com/adamhoof/ToysInterfacingBridge/pkg/TLS"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
)

func GetClientConfig() (options mqtt.ClientOptions) {
	options.AddBroker(os.Getenv("mqttServer"))
	options.SetClientID(os.Getenv("mqttClient"))
	options.SetCleanSession(false)
	options.SetOrderMatters(false)
	options.SetTLSConfig(TLS.GetConfig())
	options.SetAutoReconnect(true)
	options.SetConnectRetry(true)
	return options
}
