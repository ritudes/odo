package state

import (
	"context"

	"github.com/redhat-developer/odo/pkg/api"
)

type Client interface {
	// Init creates a devstate file for the process
	Init(ctx context.Context) error

	// SetForwardedPorts sets the forwarded ports in the state file and saves it to the file, updating the metadata
	SetForwardedPorts(ctx context.Context, fwPorts []api.ForwardedPort) error

	// GetForwardedPorts returns the ports forwarded by the current odo dev session
	GetForwardedPorts(ctx context.Context) ([]api.ForwardedPort, error)

	// SaveExit resets the state file to indicate odo is not running
	SaveExit(ctx context.Context) error
}
