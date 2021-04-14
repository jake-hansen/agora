package dto

type Invite struct {
	MeetingID string `json:"meeting_id" binding:"required"`
	MeetingPlatform string `json:"meeting_platform" binding:"required"`
	InviterID uint `json:"inviter_id" binding:"required"`
	Invitee string `json:"invitee" binding:"required"`
}

