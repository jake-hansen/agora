package meetingplatformrepomock

import (
	"github.com/stretchr/testify/mock"
)

func Provide() *MeetingPlatformRepository {
	return &MeetingPlatformRepository{mock.Mock{}}
}
