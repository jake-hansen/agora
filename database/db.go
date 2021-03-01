package database

import (
	"context"
	"fmt"
	"github.com/jake-hansen/agora/log"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Config contains the parameters for configuring a database.
type Config struct {
	URL             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	dialector       *gorm.Dialector
	Logger			*log.Log
}

// Manager manages a connection to a database.
type Manager struct {
	ctx context.Context
	cfg Config
	DB  *gorm.DB
}

// MockManager manages a mock connection to a database.
type MockManager struct {
	Manager *Manager
	Mock    *sqlmock.Sqlmock
}

// New returns a new instance of Manager configured with the given
// Config and DB.
func New(cfg Config, db *gorm.DB) *Manager {
	return &Manager{
		cfg: cfg,
		DB:  db,
	}
}

// mySQLDialector creates a Dialector neccessary for communicating
// with a MySQL database.
func mySQLDialector(v *viper.Viper) *gorm.Dialector {
	mysqlConfig := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		v.GetString("database.username"),
		v.GetString("database.password"),
		v.GetString("database.protocol"),
		v.GetString("database.host"),
		v.GetString("database.port"),
		v.GetString("database.name"))

	d := mysql.Open(mysqlConfig)
	return &d
}
