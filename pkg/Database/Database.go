package Database

import "github.com/adamhoof/ToysInterfacingBridge/pkg/Toy"

type Database interface {
	Connect()
	TestConnection()
	UpdateToyMode(toyName string, toyMode string)
	PullToyData(toyBag map[string]*Toy.Toy)
}
