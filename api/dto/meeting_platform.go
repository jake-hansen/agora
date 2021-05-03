package dto

// MeetingPlatform represents information about a meeting platform.
type MeetingPlatform struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	RedirectURL string `json:"redirect_url"`
}
