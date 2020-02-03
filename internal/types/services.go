package types

import "encoding/json"

// service representation
type Service struct {
	Name        string            `json:",omitempty"`
	Image       string            `json:",omitempty"`
	Tag         string            `json:",omitempty"`
	Labels      map[string]string `json:",omitempty"`
	Environment map[string]string `json:",omitempty"`
}

type ServiceList []Service

func (l *ServiceList) ToJSON() (string, error) {
	formatted, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "[]", err
	}
	return string(formatted), nil
}
