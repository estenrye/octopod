package types

import (
	"encoding/json"
	"github.com/docker/docker/api/types/swarm"
	"sort"
)

// Service is a simple object representation of a swarm service.
type Service struct {
	Name        string            `json:",omitempty"`
	Image       string            `json:",omitempty"`
	Tag         string            `json:",omitempty"`
	Labels      map[string]string `json:",omitempty"`
	Environment map[string]string `json:",omitempty"`
}

// ServiceList is a list of swarm services.
type ServiceList []Service

// ToJSON converts a ServiceList to its JSON representation.
func (l *ServiceList) ToJSON() (string, error) {
	formatted, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "[]", err
	}
	return string(formatted), nil
}

// TaskList provides a Sort interface for Docker Swarm Tasks
type TaskList []swarm.Task

// Less function implementation of the golang sort interface
func (t TaskList) Less(i, j int) bool {
	return t[i].ServiceID < t[j].ServiceID
}

// Swap function implementation of the golang sort interface
func (t TaskList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Len function implementation of the golang sort interface
func (t TaskList) Len() int {
	return t.Len()
}

// GetTasks returns tasks for the provided serviceID.
func (t TaskList) GetTasks(serviceID string) []swarm.Task {
	var result []swarm.Task
	sort.Sort(t)
	for _, task := range t {
		if serviceID < task.ServiceID {
			continue
		}
		if serviceID > task.ServiceID {
			break
		}
		result = append(result, task)
	}
	return result
}

// TaskSummary is a simple struct for outputing task summary information.
type TaskSummary struct {
	ID           string           `json:",omitempty"`
	Status       swarm.TaskStatus `json:",omitempty"`
	DesiredState swarm.TaskState  `json:",omitempty"`
}

// ServiceSummary is a simple struct for outputing service information with task summaries.
type ServiceSummary struct {
	Name     string            `json:",omitempty"`
	Image    string            `json:",omitempty"`
	Tag      string            `json:",omitempty"`
	Mode     swarm.ServiceMode `json:",omitempty"`
	TaskList []TaskSummary     `json:",omitempty"`
}

// ServiceSummaryList is an array of service summaries.
type ServiceSummaryList []ServiceSummary

// ToJSON converts a ServiceSummaryList to its JSON representation.
func (l *ServiceSummaryList) ToJSON() (string, error) {
	formatted, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "[]", err
	}
	return string(formatted), nil
}
