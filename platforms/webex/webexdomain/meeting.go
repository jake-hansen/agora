package webexdomain

// Meeting represents a meeting.
type Meeting struct {
	ID        string `json:"id,omitempty"`
	Title     string `json:"title"`
	Agenda    string `json:"agenda"`
	Start     string `json:"start"`
	End       string `json:"end"`
	Password  string `json:"password,omitempty"`
	HostKey   string `json:"hostKey,omitempty"`
	WebLink   string `json:"webLink,omitempty"`
	SendEmail bool   `json:"sendEmail"`
}

// MeetingList represents a list of meetings.
type MeetingList struct {
	Items []*Meeting `json:"items"`
}
