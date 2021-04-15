package dto

type MeetingPlatform struct {
	ID			uint   `json:"id"`
	Name		string `json:"name"`
	RedirectURL string	`json:"redirect_url"`
}
