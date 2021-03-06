package project

import (
	"golang.org/x/net/context"

	"github.com/docker/libcompose/config"
	"github.com/docker/libcompose/project/events"
	"github.com/docker/libcompose/project/options"
)

// APIProject defines the methods a libcompose project should implement.
type APIProject interface {
	events.Notifier
	events.Emitter

	Build(ctx context.Context, options options.Build, sevice ...string) error
	Create(ctx context.Context, options options.Create, services ...string) error
	Delete(ctx context.Context, options options.Delete, services ...string) error
	Down(ctx context.Context, options options.Down, services ...string) error
	Events(ctx context.Context, services ...string) (chan events.ContainerEvent, error)
	Kill(ctx context.Context, signal string, services ...string) error
	Log(ctx context.Context, follow bool, services ...string) error
	Pause(ctx context.Context, services ...string) error
	Ps(ctx context.Context, onlyID bool, services ...string) (InfoSet, error)
	// FIXME(vdemeester) we could use nat.Port instead ?
	Port(ctx context.Context, index int, protocol, serviceName, privatePort string) (string, error)
	Pull(ctx context.Context, services ...string) error
	Restart(ctx context.Context, timeout int, services ...string) error
	Run(ctx context.Context, serviceName string, commandParts []string) (int, error)
	Scale(ctx context.Context, timeout int, servicesScale map[string]int) error
	Start(ctx context.Context, services ...string) error
	Stop(ctx context.Context, timeout int, services ...string) error
	Unpause(ctx context.Context, services ...string) error
	Up(ctx context.Context, options options.Up, services ...string) error

	Parse() error
	CreateService(name string) (Service, error)
	AddConfig(name string, config *config.ServiceConfig) error
	Load(bytes []byte) error
}

// RuntimeProject defines runtime-specific methods for a libcompose implementation.
type RuntimeProject interface {
	RemoveOrphans(ctx context.Context, projectName string, serviceConfigs *config.ServiceConfigs) error
}
