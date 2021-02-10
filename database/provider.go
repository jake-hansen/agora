package database

import (
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Cfg provides a new Config using values from a Viper.
func Cfg(v *viper.Viper) (*Config, error) {
	c := Config{}

	if v.GetString("database.type") == "mysql" {
		c.dialector = mySQLDialector(v)
	} else {
		return nil, errors.New("database type not specified")
	}

	c.ConnMaxLifetime = v.GetDuration("database.connections.lifetime.max")
	c.MaxIdleConns = v.GetInt("database.connections.idle.max")
	c.MaxOpenConns = v.GetInt("database.connections.open.max")

	return &c, nil
}

// CfgTest provides the passed Config.
func CfgTest(cfg Config) (*Config, error) {
	return &cfg, nil
}

// ProvideGORM provides a DB using the configuration properties provided
// by the given Config.
func ProvideGORM(cfg *Config) (*gorm.DB, func(), error) {
	db, err := gorm.Open(*cfg.dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	database, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	database.SetMaxOpenConns(cfg.MaxOpenConns)
	database.SetMaxIdleConns(cfg.MaxIdleConns)
	database.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	cleanup := func() {
		if db != nil {
			_ = database.Close()
		}
	}

	return db, cleanup, err
}

// Provide provides a new Manager containing the given Config and DB.
func Provide(cfg *Config, db *gorm.DB) (*Manager, error) {
	g := New(*cfg, db)
	return g, nil
}

// ProvideMock provies a new MockManager containing the given Config.
func ProvideMock(cfg *Config) (*MockManager, func(), error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0.23"))
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if db != nil {
			_ = db.Close()
		}
	}

	g := New(*cfg, gormDB)

	manager := MockManager{
		Manager: g,
		Mock:    &mock,
	}

	return &manager, cleanup, nil
}

var (
	// ProviderProductionSet provides a new Manager for use in production.
	ProviderProductionSet = wire.NewSet(Provide, ProvideGORM, Cfg)

	// ProviderTestSet provides a new MockManager for testing.
	ProviderTestSet = wire.NewSet(ProvideMock, CfgTest)
)
