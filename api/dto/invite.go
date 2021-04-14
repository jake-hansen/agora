package dto

type Invite struct {
	MeetingID string `json:"meeting_id" binding:"required"`
	MeetingPlatform string `json:"meeting_platform" binding:"required"`
	InviterID uint `json:"inviter_id" binding:"required"`
	Invitee string `json:"invitee" binding:"required"`
}

type InviteRequest struct {
	MeetingID string `json:"meeting_id" binding:"required"`
	InviterID uint `json:"inviter_id" binding:"required"`
	InviteeUsername string `json:"invitee_username" binding:"required"`
	MeetingPlatformID uint `json:"meeting_platform_id" binding:"required"`
}

