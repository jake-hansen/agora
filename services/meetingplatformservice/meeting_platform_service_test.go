 package meetingplatformservice_test

import (
	"errors"
	"github.com/jake-hansen/agora/database/repositories/mocks/meetingplatformrepomock"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/meetingplatformservice"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var mockMeetingPlatform = domain.MeetingPlatform{
	Name:  "mock meeting provider",
}

var mockConfiguredPlatforms = []*domain.MeetingPlatform{&mockMeetingPlatform}

func TestMeetingProviderService_Save(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := meetingplatformrepomock.Provide()
		r.On("Create", mock.AnythingOfType("*domain.MeetingPlatform")).Return(1, nil)

		mService := meetingplatformservice.Provide(r, mockConfiguredPlatforms)

		id, err := mService.Save(&mockMeetingPlatform)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), id)
	})

	t.Run("failure", func(t *testing.T) {
		r := meetingplatformrepomock.Provide()
		r.On("Create", mock.AnythingOfType("*domain.MeetingPlatform")).Return(0, errors.New("unknown error"))

		mService := meetingplatformservice.Provide(r, mockConfiguredPlatforms)

		id, err := mService.Save(&mockMeetingPlatform)

		assert.Error(t, err)
		assert.Equal(t, uint(0), id)
	})
}

func TestMeetingProviderService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := meetingplatformrepomock.Provide()
		r.On("Delete", mock.AnythingOfType("uint")).Return(nil)

		mService := meetingplatformservice.Provide(r, mockConfiguredPlatforms)

		err := mService.Delete(0)

		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		r := meetingplatformrepomock.Provide()
		r.On("Delete", mock.AnythingOfType("uint")).Return(errors.New("unknown error"))

		mService := meetingplatformservice.Provide(r, mockConfiguredPlatforms)

		err := mService.Delete(0)

		assert.Error(t, err)
	})
}

func TestMeetingProviderService_GetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockProviders := []*domain.MeetingPlatform{&mockMeetingPlatform, &mockMeetingPlatform}

		r := meetingplatformrepomock.Provide()
		r.On("GetAll").Return(mockProviders, nil)

		mService := meetingplatformservice.Provide(r, mockConfiguredPlatforms)

		providers, err := mService.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, mockProviders, providers)
	})

	t.Run("failure", func(t *testing.T) {
		r := meetingplatformrepomock.Provide()
		r.On("GetAll").Return([]*domain.MeetingPlatform{}, errors.New("unknown error"))

		mService := meetingplatformservice.Provide(r, mockConfiguredPlatforms)

		providers, err := mService.GetAll()

		assert.Error(t, err)
		assert.Empty(t, providers)
	})
}

func TestMeetingProviderService_GetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := meetingplatformrepomock.Provide()
		r.On("GetByID", mock.AnythingOfType("uint")).Return(&mockMeetingPlatform, nil)

		mService := meetingplatformservice.Provide(r, mockConfiguredPlatforms)

		provider, err := mService.GetByID(0)

		assert.NoError(t, err)
		assert.Equal(t, mockMeetingPlatform, *provider)
	})

	t.Run("failure", func(t *testing.T) {
		r := meetingplatformrepomock.Provide()
		r.On("GetByID", mock.AnythingOfType("uint")).Return(&domain.MeetingPlatform{}, errors.New("unknown error"))

		mService := meetingplatformservice.Provide(r, mockConfiguredPlatforms)

		provider, err := mService.GetByID(0)

		assert.Error(t, err)
		assert.Empty(t, provider)
	})
}

func TestMeetingProviderService_GetByPlatformName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := meetingplatformrepomock.Provide()
		r.On("GetByPlatformName", mock.AnythingOfType("string")).Return(&mockMeetingPlatform, nil)

		mService := meetingplatformservice.Provide(r, mockConfiguredPlatforms)

		provider, err := mService.GetByPlatformName("test")

		assert.NoError(t, err)
		assert.Equal(t, mockMeetingPlatform, *provider)
	})

	t.Run("failure", func(t *testing.T) {
		r := meetingplatformrepomock.Provide()
		r.On("GetByPlatformName", mock.AnythingOfType("string")).Return(&domain.MeetingPlatform{}, errors.New("unknown error"))

		mService := meetingplatformservice.Provide(r, mockConfiguredPlatforms)

		provider, err := mService.GetByPlatformName("test")

		assert.Error(t, err)
		assert.Empty(t, provider)
	})
}
