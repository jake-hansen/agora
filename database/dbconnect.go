package database

import (
	"errors"
	"fmt"
	"github.com/jake-hansen/agora/config"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect connects to a MYSQL database using the parameters specified in a properties file.
// If the connection is successful, returns a pointer to gorm.DB. Otherwise, panics.
func Connect() *gorm.DB {
	c := config.GetConfig()

	if c.GetString("database.type") == "mysql" {
		return newMySQL(c)
	} else {
		panic(errors.New("database type not specified"))
	}
}

func newMySQL(config *viper.Viper) *gorm.DB {
	mysqlConfig := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetString("database.user"),
		config.GetString("database.password"),
		config.GetString("database.protocol"),
		config.GetString("database.host"),
		config.GetString("database.port"),
		config.GetString("database.name"))

	db, err := gorm.Open(mysql.Open(mysqlConfig), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
