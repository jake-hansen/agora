package domain

// Health contains information about the health of the application.
type Health struct {
	Healthy	bool
	Reason	string
}

// HealthService manages retrieving the health status of the application.
type HealthService interface {
	GetHealth() (*Health, error)
}
