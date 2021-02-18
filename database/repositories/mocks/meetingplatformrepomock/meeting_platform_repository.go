package meetingplatformrepomock

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
)

type MeetingPlatformRepository struct {
	mock.Mock
}

func (m *MeetingPlatformRepository) Create(meetingProvider *domain.MeetingPlatform) (uint, error) {
	args := m.Called(meetingProvider)
	return uint(args.Int(0)), args.Error(1)
}

func (m *MeetingPlatformRepository) GetAll() ([]*domain.MeetingPlatform, error) {
	args := m.Called()
	return args.Get(0).([]*domain.MeetingPlatform), args.Error(1)
}

func (m *MeetingPlatformRepository) GetByID(ID uint) (*domain.MeetingPlatform, error) {
	args := m.Called(ID)
	return args.Get(0).(*domain.MeetingPlatform), args.Error(1)
}

func (m *MeetingPlatformRepository) GetByProviderName(providerName string) (*domain.MeetingPlatform, error) {
	args := m.Called(providerName)
	return args.Get(0).(*domain.MeetingPlatform), args.Error(1)
}

func (m *MeetingPlatformRepository) Update(meetingProvider *domain.MeetingPlatform) error {
	args := m.Called(meetingProvider)
	return args.Error(0)
}

func (m *MeetingPlatformRepository) Delete(ID uint) error {
	args := m.Called(ID)
	return args.Error(0)
}