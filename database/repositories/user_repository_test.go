package repositories_test

import (
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
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
