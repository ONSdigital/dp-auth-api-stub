package handler

// Response represents the json body returned on requesting an identity check
type Response struct {
	Identifier string `json:"identifier,omitempty"`
	Message    string `json:"message,omitempty"`
	Name       string `json:"name,omitempty"`
	Token      string `json:"token,omitempty"`
	Status     int    `json:"-"`
}

type serviceRequest struct {
	Name string `json:"name,omitempty"`
}

// Scenarios represents a list of identities and services configured
// on application startup via a yml file
type Scenarios struct {
	Identities      []identityProfile `yaml:"identities"`
	Services        []string          `yaml:"services"`
	NewServiceToken string            `yaml:"new-service-token"`
}

type service struct {
	Name  string `yaml:"name,omitempty"`
	Token string `yaml:"token,omitempty"`
}

type identityProfile struct {
	AuthorizationToken string `yaml:"authorization-token,omitempty"`
	XFlorenceToken     string `yaml:"x-florence-token,omitempty"`
	Identifier         string `yaml:"identifier,omitempty"`
	Status             int    `yaml:"response-status,omitempty"`
	Message            string `yaml:"message,omitempty"`
	Scenario           string `yaml:"scenario,omitempty"`
}
