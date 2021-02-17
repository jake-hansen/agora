package meetingproviderrepomock

import (
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/mock"
)

type MeetingProviderRepository struct {
	mock.Mock
}

func (m *MeetingProviderRepository) Create(meetingProvider *domain.MeetingProvider) (uint, error) {
	args := m.Called(meetingProvider)
	return uint(args.Int(0)), args.Error(1)
}

func (m *MeetingProviderRepository) GetAll() ([]*domain.MeetingProvider, error) {
	args := m.Called()
	return args.Get(0).([]*domain.MeetingProvider), args.Error(1)
}

func (m *MeetingProviderRepository) GetByID(ID uint) (*domain.MeetingProvider, error) {
	args := m.Called(ID)
	return args.Get(0).(*domain.MeetingProvider), args.Error(1)
}

func (m *MeetingProviderRepository) GetByProviderName(providerName string) (*domain.MeetingProvider, error) {
	args := m.Called(providerName)
	return args.Get(0).(*domain.MeetingProvider), args.Error(1)
}

func (m *MeetingProviderRepository) Update(meetingProvider *domain.MeetingProvider) error {
	args := m.Called(meetingProvider)
	return args.Error(0)
}

func (m *MeetingProviderRepository) Delete(ID uint) error {
	args := m.Called(ID)
	return args.Error(0)
}