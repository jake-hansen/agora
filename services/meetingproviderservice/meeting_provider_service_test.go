package meetingproviderservice_test

import (
	"errors"
	"github.com/jake-hansen/agora/database/repositories/mocks/meetingproviderrepomock"
	"github.com/jake-hansen/agora/domain"
	"github.com/jake-hansen/agora/services/meetingproviderservice"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var mockMeetingProvider = domain.MeetingProvider{
	Name:  "mock meeting provider",
}

func TestMeetingProviderService_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := meetingproviderrepomock.Build()
		r.On("Create", mock.AnythingOfType("*domain.MeetingProvider")).Return(1, nil)

		mService := meetingproviderservice.BuildTest(r)

		id, err := mService.Create(&mockMeetingProvider)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), id)
	})

	t.Run("failure", func(t *testing.T) {
		r := meetingproviderrepomock.Build()
		r.On("Create", mock.AnythingOfType("*domain.MeetingProvider")).Return(0, errors.New("unknown error"))

		mService := meetingproviderservice.BuildTest(r)

		id, err := mService.Create(&mockMeetingProvider)

		assert.Error(t, err)
		assert.Equal(t, uint(0), id)
	})
}

func TestMeetingProviderService_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := meetingproviderrepomock.Build()
		r.On("Delete", mock.AnythingOfType("uint")).Return(nil)

		mService := meetingproviderservice.BuildTest(r)

		err := mService.Delete(0)

		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		r := meetingproviderrepomock.Build()
		r.On("Delete", mock.AnythingOfType("uint")).Return(errors.New("unknown error"))

		mService := meetingproviderservice.BuildTest(r)

		err := mService.Delete(0)

		assert.Error(t, err)
	})
}

func TestMeetingProviderService_GetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockProviders := []*domain.MeetingProvider{&mockMeetingProvider, &mockMeetingProvider}

		r := meetingproviderrepomock.Build()
		r.On("GetAll").Return(mockProviders, nil)

		mService := meetingproviderservice.BuildTest(r)

		providers, err := mService.GetAll()

		assert.NoError(t, err)
		assert.Equal(t, mockProviders, providers)
	})

	t.Run("failure", func(t *testing.T) {
		r := meetingproviderrepomock.Build()
		r.On("GetAll").Return([]*domain.MeetingProvider{}, errors.New("unknown error"))

		mService := meetingproviderservice.BuildTest(r)

		providers, err := mService.GetAll()

		assert.Error(t, err)
		assert.Empty(t, providers)
	})
}

func TestMeetingProviderService_GetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := meetingproviderrepomock.Build()
		r.On("GetByID", mock.AnythingOfType("uint")).Return(&mockMeetingProvider, nil)

		mService := meetingproviderservice.BuildTest(r)

		provider, err := mService.GetByID(0)

		assert.NoError(t, err)
		assert.Equal(t, mockMeetingProvider, *provider)
	})

	t.Run("failure", func(t *testing.T) {
		r := meetingproviderrepomock.Build()
		r.On("GetByID", mock.AnythingOfType("uint")).Return(&domain.MeetingProvider{}, errors.New("unknown error"))

		mService := meetingproviderservice.BuildTest(r)

		provider, err := mService.GetByID(0)

		assert.Error(t, err)
		assert.Empty(t, provider)
	})
}

func TestMeetingProviderService_GetByProviderName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := meetingproviderrepomock.Build()
		r.On("GetByProviderName", mock.AnythingOfType("string")).Return(&mockMeetingProvider, nil)

		mService := meetingproviderservice.BuildTest(r)

		provider, err := mService.GetByProviderName("test")

		assert.NoError(t, err)
		assert.Equal(t, mockMeetingProvider, *provider)
	})

	t.Run("failure", func(t *testing.T) {
		r := meetingproviderrepomock.Build()
		r.On("GetByProviderName", mock.AnythingOfType("string")).Return(&domain.MeetingProvider{}, errors.New("unknown error"))

		mService := meetingproviderservice.BuildTest(r)

		provider, err := mService.GetByProviderName("test")

		assert.Error(t, err)
		assert.Empty(t, provider)
	})
}

func TestMeetingProviderService_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := meetingproviderrepomock.Build()
		r.On("Update", mock.AnythingOfType("*domain.MeetingProvider")).Return(nil)

		mService := meetingproviderservice.BuildTest(r)

		err := mService.Update(&mockMeetingProvider)

		assert.NoError(t, err)
	})
	
	t.Run("failure", func(t *testing.T) {
		r := meetingproviderrepomock.Build()
		r.On("Update", mock.AnythingOfType("*domain.MeetingProvider")).Return(errors.New("unknown error"))

		mService := meetingproviderservice.BuildTest(r)

		err := mService.Update(&mockMeetingProvider)

		assert.Error(t, err)
	})
}
