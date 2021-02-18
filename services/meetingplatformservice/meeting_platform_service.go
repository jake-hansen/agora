package meetingplatformservice

import "github.com/jake-hansen/agora/domain"

type MeetingProviderService struct {
	repo domain.MeetingPlatformRepository
}

func (m *MeetingProviderService) Create(meetingProvider *domain.MeetingPlatform) (uint, error) {
	return m.repo.Create(meetingProvider)
}

func (m *MeetingProviderService) GetAll() ([]*domain.MeetingPlatform, error) {
	return m.repo.GetAll()
}

func (m *MeetingProviderService) GetByID(ID uint) (*domain.MeetingPlatform, error) {
	return m.repo.GetByID(ID)
}

func (m *MeetingProviderService) GetByProviderName(providerName string) (*domain.MeetingPlatform, error) {
	return m.repo.GetByProviderName(providerName)
}

func (m *MeetingProviderService) Update(meetingProvider *domain.MeetingPlatform) error {
	return m.repo.Update(meetingProvider)
}

func (m *MeetingProviderService) Delete(ID uint) error {
	return m.repo.Delete(ID)
}
