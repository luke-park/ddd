package deployment

// Payload represents a deployment payload, sent from a client to a server.
type Payload map[string]*Service

// Service represents a single service to be deployed.
type Service struct {
	Tag                  string
	EnvironmentVariables map[string]string
}
