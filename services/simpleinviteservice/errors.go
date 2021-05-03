package simpleinviteservice

type InviterSameAsInviteeErr struct{}

func NewInviterSameAsInviteeErr() InviterSameAsInviteeErr {
	return InviterSameAsInviteeErr{}
}

func (i InviterSameAsInviteeErr) Error() string {
	return "inviter cannot be the invitee"
}
