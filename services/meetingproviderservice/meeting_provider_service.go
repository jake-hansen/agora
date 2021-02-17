package meetingproviderservice

import "github.com/jake-hansen/agora/domain"

type MeetingProviderService struct {
	repo domain.MeetingProviderRepository
}

func (m *MeetingProviderService) Create(meetingProvider *domain.MeetingProvider) (uint, error) {
	return m.repo.Create(meetingProvider)
}

func (m *MeetingProviderService) GetAll() ([]*domain.MeetingProvider, error) {
	return m.repo.GetAll()
}

func (m *MeetingProviderService) GetByID(ID uint) (*domain.MeetingProvider, error) {
	return m.repo.GetByID(ID)
}

func (m *MeetingProviderService) GetByProviderName(providerName string) (*domain.MeetingProvider, error) {
	return m.repo.GetByProviderName(providerName)
}

func (m *MeetingProviderService) Update(meetingProvider *domain.MeetingProvider) error {
	return m.repo.Update(meetingProvider)
}

func (m *MeetingProviderService) Delete(ID uint) error {
	return m.repo.Delete(ID)
}
