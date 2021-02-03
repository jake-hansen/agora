package repositories_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jake-hansen/agora/database/repositories"
	"github.com/jake-hansen/agora/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
)

var mockUser = domain.User{
	Model:     gorm.Model{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: gorm.DeletedAt{},
	},
	Firstname: "john",
	Lastname:  "doe",
	Username:  "jdoe",
	Password:  "Password123!",
}

type Suite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock
	repo domain.UserRepository
}

func (s *Suite) SetupSuite() {
	db, mock, err := sqlmock.New()
	require.NoError(s.T(), err)

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.23"))

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	require.NoError(s.T(), err)

	s.DB = gormDB
	s.mock = mock
	s.repo = repositories.NewUserRepository(s.DB)
}

func (s *Suite) TestUserRepository_Create() {
	instSQL := "INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`firstname`,`lastname`,`username`," +
		"`password`) VALUES (?,?,?,?,?,?,?)"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(instSQL)).
			WithArgs(mockUser.CreatedAt, mockUser.UpdatedAt, mockUser.DeletedAt,
				mockUser.Firstname, mockUser.Lastname, mockUser.Username,
				mockUser.Password).
			WillReturnResult(sqlmock.NewResult(0, 1))
		s.mock.ExpectCommit()

		id, err := s.repo.Create(&mockUser)

		require.NoError(s.T(), err)
		assert.Equal(s.T(), uint(0), id)
	})

	s.T().Run("failure-rollback", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(instSQL)).
			WithArgs(mockUser.CreatedAt, mockUser.UpdatedAt, mockUser.DeletedAt,
				mockUser.Firstname, mockUser.Lastname, mockUser.Username,
				mockUser.Password).
			WillReturnError(errors.New("unknown error"))
		s.mock.ExpectRollback()

		id, err := s.repo.Create(&mockUser)

		require.Error(s.T(), err)
		assert.Equal(s.T(), uint(0), id)
	})
}

func (s *Suite) TestUserRepository_Delete() {
	delSQL := "UPDATE `users` SET `deleted_at`=? WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(delSQL)).
			WithArgs(sqlmock.AnyArg(), 0).
			WillReturnResult(sqlmock.NewResult(0, 1))
		s.mock.ExpectCommit()

		err := s.repo.Delete(0)

		require.NoError(t, err)
	})

	s.T().Run("failure-rollback", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(delSQL)).
			WithArgs(sqlmock.AnyArg(), 0).
			WillReturnError(errors.New("unknown error"))
		s.mock.ExpectRollback()

		err := s.repo.Delete(0)

		require.Error(t, err)
	})
}

func (s *Suite) TestUserRepository_GetByID() {
	getSQL := "SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs(0).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "firstname",
													"lastname", "username", "password"}).
				AddRow(0, mockUser.CreatedAt, mockUser.UpdatedAt, mockUser.DeletedAt, mockUser.Firstname,
					   mockUser.Lastname, mockUser.Username, mockUser.Password))

		user, err := s.repo.GetByID(0)

		require.NoError(t, err)
		assert.Equal(t, mockUser, *user)
	})

	s.T().Run("failure", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs(0).
			WillReturnError(errors.New("unknown error"))

		_, err := s.repo.GetByID(0)

		require.Error(t, err)
	})
}

func (s *Suite) TestUserRepository_GetByUsername() {
	getSQL := "SELECT * FROM `users` WHERE username = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs("jdoe").
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "firstname",
				"lastname", "username", "password"}).
				AddRow(0, mockUser.CreatedAt, mockUser.UpdatedAt, mockUser.DeletedAt, mockUser.Firstname,
					mockUser.Lastname, mockUser.Username, mockUser.Password))

		user, err := s.repo.GetByUsername("jdoe")

		require.NoError(t, err)
		assert.Equal(t, mockUser, *user)
	})

	s.T().Run("failure", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WithArgs("jdoe").
			WillReturnError(errors.New("unknown error"))

		_, err := s.repo.GetByUsername("jdoe")

		require.Error(t, err)
	})
}

func (s *Suite) TestUserRepository_GetAll() {
	getSQL := "SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "firstname",
				"lastname", "username", "password"}).
				AddRow(0, mockUser.CreatedAt, mockUser.UpdatedAt, mockUser.DeletedAt, mockUser.Firstname,
					mockUser.Lastname, mockUser.Username, mockUser.Password).
				AddRow(0, mockUser.CreatedAt, mockUser.UpdatedAt, mockUser.DeletedAt, mockUser.Firstname,
					mockUser.Lastname, mockUser.Username, mockUser.Password))

		users, err := s.repo.GetAll()

		require.NoError(t, err)
		assert.Equal(t, mockUser, *users[0])
		assert.Equal(t, mockUser, *users[1])
		assert.Len(t, users, 2)
	})

	s.T().Run("failure", func(t *testing.T) {
		s.mock.ExpectQuery(regexp.QuoteMeta(getSQL)).
			WillReturnError(errors.New("unknown error"))

		_, err := s.repo.GetAll()

		require.Error(t, err)
	})
}

func (s *Suite) TestUserRepository_Update() {
	updSQL := "UPDATE `users` SET `updated_at`=?,`firstname`=?,`lastname`=?,`username`=?,`password`=? WHERE `id` = ?"

	s.T().Run("success", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(updSQL)).
			WithArgs(sqlmock.AnyArg(), mockUser.Firstname, mockUser.Lastname, mockUser.Username, mockUser.Password, 1).
			WillReturnResult(sqlmock.NewResult(int64(mockUser.ID), 1))

		s.mock.ExpectCommit()

		user := mockUser
		user.ID = 1
		err := s.repo.Update(&user)

		require.NoError(t, err)
	})

	s.T().Run("failure-rollback", func(t *testing.T) {
		s.mock.ExpectBegin()
		s.mock.ExpectExec(regexp.QuoteMeta(updSQL)).
			WithArgs(sqlmock.AnyArg(), mockUser.Firstname, mockUser.Lastname, mockUser.Username, mockUser.Password, 1).
			WillReturnError(errors.New("unknown error"))

		s.mock.ExpectRollback()

		user := mockUser
		user.ID = 1
		err := s.repo.Update(&user)

		require.Error(t, err)
	})
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
