package loader

import (
	"github.com/jake-hansen/agora/domain"
)

// MeetingPlatformLoader provides functions for loading already configured
// MeetingPlatforms into a MeetingPlatformRepository.
type MeetingPlatformLoader struct {
	PlatformRepo        domain.MeetingPlatformRepository
	ConfiguredPlatforms domain.ConfiguredPlatforms
}

// Loader provides functions for loading data into repositories.
type Loader struct {
	meetingPlatformLoader *MeetingPlatformLoader
}

// NewLoader returns a new Loader configured with the provided MeetingPlatformLoader.
func NewLoader(meetingPlatformLoader *MeetingPlatformLoader) *Loader {
	return &Loader{
		meetingPlatformLoader: meetingPlatformLoader,
	}
}

// NewMeetingPlatformLoader returns a new MeetingPlatformLoader configured with the provided
// MeetingPlatformRepository and configured MeetingPlatforms.
func NewMeetingPlatformLoader(repository domain.MeetingPlatformRepository, configuredPlatforms []*domain.MeetingPlatform) *MeetingPlatformLoader {
	return &MeetingPlatformLoader{
		PlatformRepo:        repository,
		ConfiguredPlatforms: configuredPlatforms,
	}
}

// Load begins loading data into each configured repository.
func (l *Loader) Load() error {
	return l.meetingPlatformLoader.load()
}

// load loads the configured MeetingPlatforms into the MeetingPlatformRepository.
// If a MeetingPlatform already exists (matching by name only) then that platform
// is not loaded again.
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
