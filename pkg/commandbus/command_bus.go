package commandbus

import (
	"context"

	"github.com/vardius/go-api-boilerplate/pkg/domain"
)

// CommandHandler function
type CommandHandler interface{}

// CommandBus allows to subscribe/dispatch commands
type CommandBus interface {
	Publish(ctx context.Context, command domain.Command, out chan<- error)
	Subscribe(commandName string, fn CommandHandler) error
	Unsubscribe(commandName string, fn CommandHandler) error
}
