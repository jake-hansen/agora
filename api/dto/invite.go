package dto

type Invite struct {
	ID			      uint `json:"id"`
	MeetingPlatformID uint `json:"meeting_platform" binding:"required"`
	InviterID uint `json:"inviter_id" binding:"required"`
	InviteeID uint `json:"invitee" binding:"required"`
	Meeting Meeting `json:"meeting"`
}

type InviteRequest struct {
	MeetingID string `json:"meeting_id" binding:"required"`
	InviterID uint `json:"inviter_id" binding:"required"`
	InviteeUsername string `json:"invitee_username" binding:"required"`
	MeetingPlatformID uint `json:"meeting_platform_id" binding:"required"`
}

