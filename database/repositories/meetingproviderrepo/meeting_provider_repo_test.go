package meetingproviderrepo_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/repositories/meetingproviderrepo"
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

var mockMeetingProvider = domain.MeetingProvider {
	Model: gorm.Model{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	},
	Name:  "really awesome meeting provider",
}

type Suite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo domain.MeetingProviderRepository
}

func (s *Suite) SetupSuite()  {
	manager, _, err := database.BuildTest(database.Config{})
	s.Require().NoError(err)

	s.mock = *manager.Mock
	s.repo, _, err = meetingproviderrepo.Build(manager.Manager)
	s.Require().NoError(err)
}

func (s *Suite) TestMeetingProviderRepo_Create() {
	instSQL := "INSERT INTO `meeting_providers` (`created_at`,`updated_at`,`deleted_at`,`name`) VALUES (?,?,?,?)"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(instSQL)).
			WithArgs(mockMeetingProvider.CreatedAt, mockMeetingProvider.UpdatedAt,
				mockMeetingProvider.DeletedAt, mockMeetingProvider.Name).
			WillReturnResult(sqlmock.NewResult(0, 1))
		s.mock.ExpectCommit()

		id, err := s.repo.Create(&mockMeetingProvider)

		require.NoError(t, err)
		assert.Equal(t, uint(0), id)
	})

	s.T().Run("failure-rollback", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(instSQL)).
			WithArgs(mockMeetingProvider.CreatedAt, mockMeetingProvider.UpdatedAt,
				mockMeetingProvider.DeletedAt, mockMeetingProvider.Name).
			WillReturnError(errors.New("unknown error"))
		s.mock.ExpectRollback()

		_, err := s.repo.Create(&mockMeetingProvider)
		require.Error(t, err)
	})
}

func (s *Suite) TestMeetingProviderRepo_Delete() {
	delSQL := "UPDATE `meeting_providers` SET `deleted_at`=? WHERE `meeting_providers`.`id` = ? " +
		"AND `meeting_providers`.`deleted_at` IS NULL"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(delSQL)).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		s.mock.ExpectCommit()

		err := s.repo.Delete(1)

		require.NoError(t, err)
	})

	s.T().Run("failure-rollback", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(delSQL)).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnError(errors.New("unknown error"))
		s.mock.ExpectRollback()

		err := s.repo.Delete(1)

		require.Error(t, err)
	})
}

func (s *Suite) TestMeetingProviderRepo_GetAll() {
	getSQL := "SELECT * FROM `meeting_providers` WHERE `meeting_providers`.`deleted_at` IS NULL"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name"}).
				AddRow(0, mockMeetingProvider.CreatedAt, mockMeetingProvider.UpdatedAt,
					mockMeetingProvider.DeletedAt, mockMeetingProvider.Name).
				AddRow(0, mockMeetingProvider.CreatedAt, mockMeetingProvider.UpdatedAt,
					mockMeetingProvider.DeletedAt, mockMeetingProvider.Name))

		providers, err := s.repo.GetAll()
		require.NoError(t, err)

		assert.Equal(t, &mockMeetingProvider, providers[0])
		assert.Equal(t, &mockMeetingProvider, providers[1])
		assert.Len(t, providers, 2)
	})

	s.T().Run("failure", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WillReturnError(errors.New("unknown error"))

		_, err := s.repo.GetAll()

		require.Error(t, err)
	})
}

func (s *Suite)TestMeetingProviderRepo_GetByID() {
	getSQL := "SELECT * FROM `meeting_providers` WHERE `meeting_providers`.`id` = ? AND " +
		"`meeting_providers`.`deleted_at` IS NULL ORDER BY `meeting_providers`.`id` LIMIT 1"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs(0).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name"}).
				AddRow(0, mockMeetingProvider.CreatedAt, mockMeetingProvider.UpdatedAt,
					mockMeetingProvider.DeletedAt, mockMeetingProvider.Name))

		provider, err := s.repo.GetByID(0)

		require.NoError(t, err)
		assert.Equal(t, &mockMeetingProvider, provider)
	})

	s.T().Run("failure", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs(0).
			WillReturnError(errors.New("unknown error"))

		_, err := s.repo.GetByID(0)

		require.Error(t, err)
	})
}

func (s *Suite) TestMeetingProviderRepo_GetByProviderName() {
	getSQL := "SELECT * FROM `meeting_providers` WHERE name = ? AND " +
		"`meeting_providers`.`deleted_at` IS NULL ORDER BY `meeting_providers`.`id` LIMIT 1"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs(mockMeetingProvider.Name).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name"}).
				AddRow(0, mockMeetingProvider.CreatedAt, mockMeetingProvider.UpdatedAt,
					mockMeetingProvider.DeletedAt, mockMeetingProvider.Name))

		provider, err := s.repo.GetByProviderName(mockMeetingProvider.Name)

		require.NoError(t, err)
		assert.Equal(t, &mockMeetingProvider, provider)
	})

	s.T().Run("failure", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs(mockMeetingProvider.Name).
			WillReturnError(errors.New("unknown error"))

		_, err := s.repo.GetByProviderName(mockMeetingProvider.Name)

		require.Error(t, err)
	})
}

func (s *Suite) TestMeetingProviderRepo_Update() {
	updSQL := "UPDATE `meeting_providers` SET `updated_at`=?,`name`=? WHERE `id` = ?"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(updSQL)).
			WithArgs(sqlmock.AnyArg(), mockMeetingProvider.Name, 1).
			WillReturnResult(sqlmock.NewResult(int64(mockMeetingProvider.ID), 1))
		s.mock.ExpectCommit()

		provider := mockMeetingProvider
		provider.ID = 1
		err := s.repo.Update(&provider)

		require.NoError(t, err)
	})

	s.T().Run("failure-rollback", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(updSQL)).
			WithArgs(sqlmock.AnyArg(), mockMeetingProvider.Name, 1).
			WillReturnError(errors.New("unknown error"))
		s.mock.ExpectRollback()

		err := s.repo.Update(&mockMeetingProvider)

		require.Error(t, err)
	})
}

func TestMeetingProviderRepoSuite(t *testing.T)  {
	suite.Run(t, new(Suite))
}