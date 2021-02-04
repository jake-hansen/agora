package database

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Config struct {
	URL             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	dialector		*gorm.Dialector
}

type Manager struct {
	ctx context.Context
	cfg Config
	DB  *gorm.DB
}

type MockManager struct {
	Manager
	Mock *sqlmock.Sqlmock
}

func New(cfg Config, db *gorm.DB) *Manager {
	return &Manager{
		cfg: cfg,
		DB:  db,
	}
}

func mySQLDialector(v *viper.Viper) *gorm.Dialector {
	mysqlConfig := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		v.GetString("database.user"),
		v.GetString("database.password"),
		v.GetString("database.protocol"),
		v.GetString("database.host"),
		v.GetString("database.port"),
		v.GetString("database.name"))

	d := mysql.Open(mysqlConfig)
	return &d
}
