package types

// service representation
type Service struct {
	Name        string            `json:",omitempty"`
	Image       string            `json:",omitempty"`
	Tag         string            `json:",omitempty"`
	Labels      map[string]string `json:",omitempty"`
	Environment []string          `json:",omitempty"`
}
