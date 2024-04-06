package models

import "encoding/json"

type Node struct {
	ID         string      `json:"id"`
	Containers []Container `json:"containers"`
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