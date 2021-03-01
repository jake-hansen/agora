package domain

type Health struct {
	Healthy	bool
	Reason	string
}

type HealthService interface {
	GetHealth() (*Health, error)
}
