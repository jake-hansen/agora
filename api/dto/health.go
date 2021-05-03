package dto

// Health represents health status of the application.
type Health struct {
	Healthy bool   `json:"healthy"`
	Reason  string `json:"reason,omitempty"`
}
