package types

import (
	"encoding/json"
	"github.com/docker/docker/api/types/swarm"
)

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

// Sort interface for Docker Swarm Tasks
type TaskList []swarm.Task

// Implementation of the Less function of the golang sort interface
func (t TaskList) Less(i, j int) bool {
	return t[i].ServiceID < t[j].ServiceID && t[i].ID < t[j].ID
}

// Implementation of the Swap function of the golang sort interface
func (t TaskList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Implementation of the Len function of the golang sort interface
func (t TaskList) Len() int {
	return t.Len()
}

type TaskSummary struct {
	Id           string           `json:",omitempty"`
	Status       swarm.TaskStatus `json:",omitempty"`
	DesiredState swarm.TaskState  `json:",omitempty"`
}

type ServiceSummary struct {
	Name     string            `json:",omitempty"`
	Image    string            `json:",omitempty"`
	Tag      string            `json:",omitempty"`
	Mode     swarm.ServiceMode `json:",omitempty"`
	TaskList []TaskSummary     `json:",omitempty"`
}

type ServiceSummaryList []ServiceSummary

func (l *ServiceSummaryList) ToJSON() (string, error) {
	formatted, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "[]", err
	}
	return string(formatted), nil
}
