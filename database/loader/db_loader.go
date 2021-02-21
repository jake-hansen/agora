package loader

import "github.com/jake-hansen/agora/domain"

type MeetingPlatformLoader struct {
	PlatformRepo	domain.MeetingPlatformRepository
	ConfiguredPlatforms	[]*domain.MeetingPlatform
}

type Loader struct {
	meetingPlatformLoader *MeetingPlatformLoader
}

func NewLoader(meetingPlatformLoader *MeetingPlatformLoader) *Loader {
	return &Loader{
		meetingPlatformLoader: meetingPlatformLoader,
	}
}

func NewMeetingPlatformLoader(repository domain.MeetingPlatformRepository, configuredPlatforms []*domain.MeetingPlatform) *MeetingPlatformLoader {
	return &MeetingPlatformLoader{
		PlatformRepo:        repository,
		ConfiguredPlatforms: configuredPlatforms,
	}
}

func (l *Loader) Load() error {
	return l.meetingPlatformLoader.load()
}

func (m *MeetingPlatformLoader) load() error {
	platforms, err := m.PlatformRepo.GetAll()
	if err != nil {
		return err
	}

	platformMap := make(map[string]*domain.MeetingPlatform)

	for _, dbPlatform := range platforms {
		platformMap[dbPlatform.Name] = dbPlatform
	}

	for _, configuredPlatform := range m.ConfiguredPlatforms {
		if platformMap[configuredPlatform.Name] == nil {
			_, err = m.PlatformRepo.Create(configuredPlatform)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
