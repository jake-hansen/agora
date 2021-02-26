package dto

type Health struct {
	Healthy bool `json:"healthy"`
	Reason	string `json:"reason,omitempty"`
}
