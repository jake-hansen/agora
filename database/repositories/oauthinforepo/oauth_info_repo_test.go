package oauthinforepo_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jake-hansen/agora/database"
	"github.com/jake-hansen/agora/database/repositories/oauthinforepo"
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

var mockOAuthInfo = domain.OAuthInfo{
	Model:             gorm.Model{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	},
	UserID:            uint(1),
	MeetingPlatformID: uint(123456),
	AccessToken:       "random-access-token",
	RefreshToken:      "random-refresh-token",
}

type Suite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	repo domain.OAuthInfoRepository
}

func (s *Suite) SetupTest() {
	manager, _, err := database.BuildTest(database.Config{})
	s.Require().NoError(err)

	s.mock = *manager.Mock
	s.repo, _, err = oauthinforepo.Build(manager.Manager)
	s.Require().NoError(err)
}

func (s *Suite) TestOAuthInfoRepo_Create() {
	instSQL := "INSERT INTO `oauth_info` (`created_at`,`updated_at`,`deleted_at`,`user_id`,`meeting_provider_id`,`access_token`,`refresh_token`) VALUES (?,?,?,?,?,?,?)"

	s.T().Run("success", func(t *testing.T) {
		defer assert.NoError(t, s.mock.ExpectationsWereMet())
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(instSQL)).
			WithArgs(mockOAuthInfo.CreatedAt, mockOAuthInfo.UpdatedAt, mockOAuthInfo.DeletedAt, mockOAuthInfo.UserID, mockOAuthInfo.MeetingPlatformID, mockOAuthInfo.AccessToken, mockOAuthInfo.RefreshToken).
			WillReturnResult(sqlmock.NewResult(0, 1))

		s.mock.ExpectCommit()

		id, err := s.repo.Create(&mockOAuthInfo)

		require.NoError(t, err)
		assert.Equal(t, uint(0), id)
	})

	s.T().Run("failure-rollback", func(t *testing.T) {
		defer assert.NoError(t, s.mock.ExpectationsWereMet())
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(instSQL)).
			WithArgs(mockOAuthInfo.CreatedAt, mockOAuthInfo.UpdatedAt, mockOAuthInfo.DeletedAt, mockOAuthInfo.UserID, mockOAuthInfo.MeetingPlatformID, mockOAuthInfo.AccessToken, mockOAuthInfo.RefreshToken).
			WillReturnError(errors.New("unknown error"))
		s.mock.ExpectRollback()

		_, err := s.repo.Create(&mockOAuthInfo)
		require.Error(t, err)
	})
}

func (s *Suite) TestOAuthInfoRepo_Delete() {
	delSQL := "UPDATE `oauth_info` SET `deleted_at`=? WHERE `oauth_info`.`id` = ? AND `oauth_info`.`deleted_at` IS NULL"

	s.T().Run("success", func(t *testing.T) {
		defer assert.NoError(t, s.mock.ExpectationsWereMet())
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(delSQL)).
			WithArgs(sqlmock.AnyArg(), 0).
			WillReturnResult(sqlmock.NewResult(0, 1))
		s.mock.ExpectCommit()

		err := s.repo.Delete(0)

		require.NoError(t, err)
	})

	s.T().Run("failure-rollback", func(t *testing.T) {
		defer assert.NoError(t, s.mock.ExpectationsWereMet())
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(delSQL)).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnError(errors.New("unknown error"))
		s.mock.ExpectRollback()

		err := s.repo.Delete(1)

		require.Error(t, err)
	})
}

func (s *Suite) TestOAuthInfoRepo_GetAll() {
	getSQL := "SELECT * FROM `oauth_info` WHERE `oauth_info`.`deleted_at` IS NULL"

	s.T().Run("success", func(t *testing.T) {
		defer assert.NoError(t, s.mock.ExpectationsWereMet())
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "user_id", "access_token", "refresh_token", "meeting_provider_id"}).
				AddRow(0, mockOAuthInfo.CreatedAt, mockOAuthInfo.UpdatedAt, mockOAuthInfo.DeletedAt, mockOAuthInfo.UserID, mockOAuthInfo.AccessToken, mockOAuthInfo.RefreshToken, mockOAuthInfo.MeetingPlatformID).
				AddRow(0, mockOAuthInfo.CreatedAt, mockOAuthInfo.UpdatedAt, mockOAuthInfo.DeletedAt, mockOAuthInfo.UserID, mockOAuthInfo.AccessToken, mockOAuthInfo.RefreshToken, mockOAuthInfo.MeetingPlatformID))

		oauthinfos, err := s.repo.GetAll()
		require.NoError(t, err)

		assert.Equal(t, &mockOAuthInfo, oauthinfos[0])
		assert.Equal(t, &mockOAuthInfo, oauthinfos[1])

		assert.Len(t, oauthinfos, 2)
	})

	s.T().Run("failure", func(t *testing.T) {
		defer assert.NoError(t, s.mock.ExpectationsWereMet())
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WillReturnError(errors.New("unknown error"))

		_, err := s.repo.GetAll()

		require.Error(t, err)
	})
}

func (s *Suite) TestOAuthInfoRepo_GetAllByMeetingProviderId() {
	getSQL := "SELECT * FROM `oauth_info` WHERE meeting_provider_id = ? AND `oauth_info`.`deleted_at` IS NULL"

	s.T().Run("success", func(t *testing.T) {
		defer assert.NoError(t, s.mock.ExpectationsWereMet())
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs(mockOAuthInfo.MeetingPlatformID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "user_id", "access_token", "refresh_token", "meeting_provider_id"}).
				AddRow(0, mockOAuthInfo.CreatedAt, mockOAuthInfo.UpdatedAt, mockOAuthInfo.DeletedAt, mockOAuthInfo.UserID, mockOAuthInfo.AccessToken, mockOAuthInfo.RefreshToken, mockOAuthInfo.MeetingPlatformID).
				AddRow(0, mockOAuthInfo.CreatedAt, mockOAuthInfo.UpdatedAt, mockOAuthInfo.DeletedAt, mockOAuthInfo.UserID, mockOAuthInfo.AccessToken, mockOAuthInfo.RefreshToken, mockOAuthInfo.MeetingPlatformID))

		oauthinfos, err := s.repo.GetAllByMeetingProviderId(mockOAuthInfo.MeetingPlatformID)
		require.NoError(t, err)

		assert.Equal(t, &mockOAuthInfo, oauthinfos[0])
		assert.Equal(t, &mockOAuthInfo, oauthinfos[1])

		assert.Len(t, oauthinfos, 2)
	})

	s.T().Run("failure", func(t *testing.T) {
		defer assert.NoError(t, s.mock.ExpectationsWereMet())
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs(mockOAuthInfo.MeetingPlatformID).
			WillReturnError(errors.New("unknown error"))

		_, err := s.repo.GetAllByMeetingProviderId(mockOAuthInfo.MeetingPlatformID)

		require.Error(t, err)
	})
}

func (s *Suite) TestOAuthInfoRepo_GetAllByUserID() {
	getSQL := "SELECT * FROM `oauth_info` WHERE user_id = ? AND `oauth_info`.`deleted_at` IS NULL"

	s.T().Run("success", func(t *testing.T) {
		defer assert.NoError(t, s.mock.ExpectationsWereMet())
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs(mockOAuthInfo.UserID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "user_id", "access_token", "refresh_token", "meeting_provider_id"}).
				AddRow(0, mockOAuthInfo.CreatedAt, mockOAuthInfo.UpdatedAt, mockOAuthInfo.DeletedAt, mockOAuthInfo.UserID, mockOAuthInfo.AccessToken, mockOAuthInfo.RefreshToken, mockOAuthInfo.MeetingPlatformID).
				AddRow(0, mockOAuthInfo.CreatedAt, mockOAuthInfo.UpdatedAt, mockOAuthInfo.DeletedAt, mockOAuthInfo.UserID, mockOAuthInfo.AccessToken, mockOAuthInfo.RefreshToken, mockOAuthInfo.MeetingPlatformID))

		oauthinfos, err := s.repo.GetAllByUserID(mockOAuthInfo.UserID)
		require.NoError(t, err)

		assert.Equal(t, &mockOAuthInfo, oauthinfos[0])
		assert.Equal(t, &mockOAuthInfo, oauthinfos[1])

		assert.Len(t, oauthinfos, 2)
	})

	s.T().Run("failure", func(t *testing.T) {
		defer assert.NoError(t, s.mock.ExpectationsWereMet())
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs(mockOAuthInfo.UserID).
			WillReturnError(errors.New("unknown error"))

		_, err := s.repo.GetAllByUserID(mockOAuthInfo.UserID)

		require.Error(t, err)
	})
}

func (s *Suite) TestOAuthInfoRepo_GetByID() {
	getSQL := "SELECT * FROM `oauth_info` WHERE `oauth_info`.`id` = ? AND `oauth_info`.`deleted_at` IS NULL ORDER BY `oauth_info`.`id` LIMIT 1"

	s.T().Run("success", func(t *testing.T) {
		defer assert.NoError(t, s.mock.ExpectationsWereMet())
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs(0).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "user_id", "access_token", "refresh_token", "meeting_provider_id"}).
				AddRow(0, mockOAuthInfo.CreatedAt, mockOAuthInfo.UpdatedAt, mockOAuthInfo.DeletedAt, mockOAuthInfo.UserID, mockOAuthInfo.AccessToken, mockOAuthInfo.RefreshToken, mockOAuthInfo.MeetingPlatformID))

		oauthinfo, err := s.repo.GetByID(0)

		require.NoError(t, err)
		assert.Equal(t, &mockOAuthInfo, oauthinfo)
})

	s.T().Run("failure", func(t *testing.T) {
		defer assert.NoError(t, s.mock.ExpectationsWereMet())
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs(0).
			WillReturnError(errors.New("unknown error"))

		_, err := s.repo.GetByID(0)

		require.Error(t, err)
	})
}

func (s *Suite) TestOAuthInfoRepo_Update() {
	updSQL := "UPDATE `oauth_info` SET `updated_at`=?,`user_id`=?,`meeting_provider_id`=?,`access_token`=?,`refresh_token`=? WHERE `id` = ?"

	s.T().Run("success", func(t *testing.T) {
		defer assert.NoError(t, s.mock.ExpectationsWereMet())
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(updSQL)).
			WithArgs(sqlmock.AnyArg(), mockOAuthInfo.UserID, mockOAuthInfo.MeetingPlatformID, mockOAuthInfo.AccessToken, mockOAuthInfo.RefreshToken, 1).
			WillReturnResult(sqlmock.NewResult(int64(mockOAuthInfo.ID), 1))
		s.mock.ExpectCommit()

		oauth := mockOAuthInfo
		oauth.ID = 1
		err := s.repo.Update(&oauth)

		require.NoError(t, err)
	})

	s.T().Run("failure-rollback", func(t *testing.T) {
		defer assert.NoError(t, s.mock.ExpectationsWereMet())
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(updSQL)).
			WithArgs(sqlmock.AnyArg(), mockOAuthInfo.UserID, mockOAuthInfo.MeetingPlatformID, mockOAuthInfo.AccessToken, mockOAuthInfo.RefreshToken, 1).
			WillReturnError(errors.New("unknown error"))
		s.mock.ExpectRollback()

		oauth := mockOAuthInfo
		oauth.ID = 1
		err := s.repo.Update(&oauth)

		require.Error(t, err)
	})
}

func TestOauthInfoRepoSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}