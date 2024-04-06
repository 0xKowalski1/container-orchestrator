package statemanager

import "encoding/json"

// Node - /nodes
type Node struct {
	ID           string      `json:"id"`
	ContainerIDs []string    `json:"containerIDs"`
	Containers   []Container `json:"containers,omitempty"`
}

func (n Node) Key() string {
	return "/nodes/" + n.ID
}

func (n Node) Value() (string, error) {
	bytes, err := json.Marshal(n)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Namespace - /namespaces
type Namespace struct {
	ID string // Id is namespace value
}

func (n Namespace) Key() string {
	return "/namespaces/" + n.ID
}

func (n Namespace) Value() (string, error) {
	bytes, err := json.Marshal(n)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Container - /namespaces/{namespace}/containers
type Container struct {
	ID            string
	DesiredStatus string // running or stopped
	Status        string
	NamespaceID   string
	NodeID        string
	Image         string
	Env           []string
	StopTimeout   int
}

type CreateContainerRequest struct {
	ID          string   `json:"id"`
	Image       string   `json:"image"`
	Env         []string `json:"env"`
	StopTimeout int      `json:"stopTimeout"`
}

type UpdateContainerRequest struct {
	DesiredStatus *string `json:"desiredStatus,omitempty"` // Pointer allows differentiation between an omitted field and an empty value
	NodeID        *string `json:"nodeId,omitempty"`
	Status        *string `json:"status,omitempty"`
}

func (c Container) Key() string {
	return "/namespaces/" + c.NamespaceID + "/containers/" + c.ID
}

func (c Container) Value() (string, error) {
	bytes, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
