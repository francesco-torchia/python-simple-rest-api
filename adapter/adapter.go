package adapter

import (
	"github.com/torchiaf/Sensors/adapter/config"
	"github.com/torchiaf/Sensors/adapter/models"
	"github.com/torchiaf/Sensors/adapter/rabbitmq"
)

func Read(module string, device string, args any) (any, error) {
	routingKey := config.Config.Modules[module].RoutingKey

	rabbitmq.Exec(
		routingKey,
		models.Message{
			Device: device,
			Action: "read",
		},
	)

	return nil, nil
}

func Write(module string, device string, args any) (any, error) {
	routingKey := config.Config.Modules[module].RoutingKey

	rabbitmq.Exec(
		routingKey,
		models.Message{
			Device: device,
			Action: "write",
		},
	)

	return nil, nil
}
