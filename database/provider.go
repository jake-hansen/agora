package database

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Cfg(v *viper.Viper) (*Config, error) {
	c := Config{}

	if v.GetString("database.type") == "mysql" {
		c.dialector = mySQLDialector(v)
	} else {
		return nil, errors.New("database type not specified")
	}

	c.ConnMaxLifetime = v.GetDuration("database.connections.lifetime.max")
	c.MaxIdleConns	  = v.GetInt("database.connections.idle.max")
	c.MaxOpenConns	  = v.GetInt("database.connections.open.max")

	return &c, nil
}

func CfgTest(cfg Config) (*Config, error) {
	return &cfg, nil
}

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

func ProvideGORMTest() (*gorm.DB, func(), error) {
	db, _, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if db != nil {
			_ = db.Close()
		}
	}

	return gormDB, cleanup, nil
}

func Provide(cfg *Config, db *gorm.DB) (*Manager, error) {
	g := New(*cfg, db)
	return g, nil
}

var (
	ProviderProductionSet = wire.NewSet(Provide, ProvideGORM, Cfg)
	ProviderTestSet       = wire.NewSet(Provide, ProvideGORMTest, CfgTest)
)
