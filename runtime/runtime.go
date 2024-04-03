package runtime

import (
	"fmt"
)

// NewRuntime creates and returns a Runtime implementation based on the provided config.
func NewRuntime(backend string) (Runtime, error) {
	switch backend {
	case "containerd":
		return NewContainerdRuntime("/run/containerd/containerd.sock")
	default:
		return nil, fmt.Errorf("unsupported runtime type: %s", backend)
	}
}

// Runtime defines the interface for a container runtime.
type Runtime interface {
	// CreateContainer instantiates a new container but does not start it.
	//CreateContainer(ctx context.Context, config ContainerConfig) (Container, error)

	// StartContainer starts an existing container.
	//StartContainer(ctx context.Context, containerID string) error

	// StopContainer stops a running container.
	//StopContainer(ctx context.Context, containerID string, timeout int) error

	// RemoveContainer removes a container from the system. This may require the container to be stopped first.
	//RemoveContainer(ctx context.Context, containerID string) error

	// ListContainers returns a list of all containers managed by the runtime.
	ListContainers(namespace string) ([]Container, error)

	// GetContainerLogs returns the logs for a specific container.
	//GetContainerLogs(ctx context.Context, containerID string) (io.ReadCloser, error)

	// InspectContainer returns detailed information about a specific container.
	//InspectContainer(ctx context.Context, containerID string) (Container, error)
}

type ContainerConfig struct {
	Image string            // The container image to use.
	Cmd   []string          // Command to run in the container.
	Env   map[string]string // Environment variables for the container.
	Ports []int             // Ports to expose from the container.
}

type Container struct {
	ID      string // Unique identifier for the container.
	Status  string // Current status of the container (e.g., "running", "stopped").
	Address string // Network address for accessing the container, if applicable.
}