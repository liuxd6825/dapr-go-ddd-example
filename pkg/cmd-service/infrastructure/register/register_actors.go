package register

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
	"github.com/liuxd6825/dapr-go-sdk/actor"
)

func GetActors() []actor.FactoryContext {
	return restapp.GetActors()
}
