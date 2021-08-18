package helpers

import (
	"awesomebook/models"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var sugar = GetSugar()

func LoadDBConfig() (models.DBConfig, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	var dBConfig models.DBConfig
	if err := viper.ReadInConfig(); err != nil {
		return dBConfig, err
	}

	err := viper.Unmarshal(&dBConfig)

	if err != nil {
		return dBConfig, err
	}

	sugar.Infof("==> Get database config done: %s", dBConfig.String())
	return dBConfig, nil
}

func DBConn() (db *gorm.DB) {
	dBConfig, err := LoadDBConfig()
	if err != nil {
		panic(err.Error())
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dBConfig.Username, dBConfig.Passwd, dBConfig.Host, dBConfig.Port, dBConfig.DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		panic(err.Error())
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	sugar.Infof("==> Connected databse: %s", dsn)

	return
}
