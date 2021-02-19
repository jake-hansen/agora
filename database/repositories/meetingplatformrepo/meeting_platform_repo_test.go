package meetingplatformrepo_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/repositories/meetingplatformrepo"
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

var mockMeetingPlatform = domain.MeetingPlatform{
	Model: gorm.Model{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	},
	Name:  "really awesome meeting platform",
	RedirectURL: "https://my-redirect",
}

type Suite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo domain.MeetingPlatformRepository
}

func (s *Suite) SetupTest()  {
	manager, _, err := database.BuildTest(database.Config{})
	s.Require().NoError(err)

	s.mock = *manager.Mock
	s.repo, _, err = meetingplatformrepo.Build(manager.Manager)
	s.Require().NoError(err)
}

func (s *Suite) TestMeetingProviderRepo_Create() {
	instSQL := "INSERT INTO `meeting_platforms` (`created_at`,`updated_at`,`deleted_at`,`name`,`redirect_url`) VALUES (?,?,?,?,?)"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(instSQL)).
			WithArgs(mockMeetingPlatform.CreatedAt, mockMeetingPlatform.UpdatedAt,
				mockMeetingPlatform.DeletedAt, mockMeetingPlatform.Name, mockMeetingPlatform.RedirectURL).
			WillReturnResult(sqlmock.NewResult(0, 1))
		s.mock.ExpectCommit()

		id, err := s.repo.Create(&mockMeetingPlatform)

		require.NoError(t, err)
		assert.Equal(t, uint(0), id)
	})

	s.T().Run("failure-rollback", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(instSQL)).
			WithArgs(mockMeetingPlatform.CreatedAt, mockMeetingPlatform.UpdatedAt,
				mockMeetingPlatform.DeletedAt, mockMeetingPlatform.Name, mockMeetingPlatform.RedirectURL).
			WillReturnError(errors.New("unknown error"))
		s.mock.ExpectRollback()

		_, err := s.repo.Create(&mockMeetingPlatform)
		require.Error(t, err)
	})
}

func (s *Suite) TestMeetingProviderRepo_Delete() {
	delSQL := "UPDATE `meeting_platforms` SET `deleted_at`=? WHERE `meeting_platforms`.`id` = ? " +
		"AND `meeting_platforms`.`deleted_at` IS NULL"

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
	getSQL := "SELECT * FROM `meeting_platforms` WHERE `meeting_platforms`.`deleted_at` IS NULL"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "redirect_url"}).
				AddRow(0, mockMeetingPlatform.CreatedAt, mockMeetingPlatform.UpdatedAt,
					mockMeetingPlatform.DeletedAt, mockMeetingPlatform.Name, mockMeetingPlatform.RedirectURL).
				AddRow(0, mockMeetingPlatform.CreatedAt, mockMeetingPlatform.UpdatedAt,
					mockMeetingPlatform.DeletedAt, mockMeetingPlatform.Name, mockMeetingPlatform.RedirectURL))

		providers, err := s.repo.GetAll()
		require.NoError(t, err)

		assert.Equal(t, &mockMeetingPlatform, providers[0])
		assert.Equal(t, &mockMeetingPlatform, providers[1])
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
	getSQL := "SELECT * FROM `meeting_platforms` WHERE `meeting_platforms`.`id` = ? AND " +
		"`meeting_platforms`.`deleted_at` IS NULL ORDER BY `meeting_platforms`.`id` LIMIT 1"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs(0).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "redirect_url"}).
				AddRow(0, mockMeetingPlatform.CreatedAt, mockMeetingPlatform.UpdatedAt,
					mockMeetingPlatform.DeletedAt, mockMeetingPlatform.Name, mockMeetingPlatform.RedirectURL))

		provider, err := s.repo.GetByID(0)

		require.NoError(t, err)
		assert.Equal(t, &mockMeetingPlatform, provider)
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
	getSQL := "SELECT * FROM `meeting_platforms` WHERE name = ? AND " +
		"`meeting_platforms`.`deleted_at` IS NULL ORDER BY `meeting_platforms`.`id` LIMIT 1"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs(mockMeetingPlatform.Name).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "name", "redirect_url"}).
				AddRow(0, mockMeetingPlatform.CreatedAt, mockMeetingPlatform.UpdatedAt,
					mockMeetingPlatform.DeletedAt, mockMeetingPlatform.Name, mockMeetingPlatform.RedirectURL))

		provider, err := s.repo.GetByPlatformName(mockMeetingPlatform.Name)

		require.NoError(t, err)
		assert.Equal(t, &mockMeetingPlatform, provider)
	})

	s.T().Run("failure", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs(mockMeetingPlatform.Name).
			WillReturnError(errors.New("unknown error"))

		_, err := s.repo.GetByPlatformName(mockMeetingPlatform.Name)

		require.Error(t, err)
	})
}

func (s *Suite) TestMeetingProviderRepo_Update() {
	updSQL := "UPDATE `meeting_platforms` SET `updated_at`=?,`name`=? WHERE `id` = ?"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(updSQL)).
			WithArgs(sqlmock.AnyArg(), mockMeetingPlatform.Name, 1).
			WillReturnResult(sqlmock.NewResult(int64(mockMeetingPlatform.ID), 1))
		s.mock.ExpectCommit()

		provider := mockMeetingPlatform
		provider.ID = 1
		err := s.repo.Update(&provider)

		require.NoError(t, err)
	})

	s.T().Run("failure-rollback", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(updSQL)).
			WithArgs(sqlmock.AnyArg(), mockMeetingPlatform.Name, 1).
			WillReturnError(errors.New("unknown error"))
		s.mock.ExpectRollback()

		provider := mockMeetingPlatform
		provider.ID = 1
		err := s.repo.Update(&provider)

		require.Error(t, err)
	})
}

func TestMeetingProviderRepoSuite(t *testing.T)  {
	suite.Run(t, new(Suite))
}