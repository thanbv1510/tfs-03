package database

import (
	"context"
	"csv-import-mysql/models"
	"csv-import-mysql/utils"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"time"
)

var sugar = utils.GetSugar()

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

	//sugar.Infof("==> Get database config done: %s", dBConfig.String())
	return dBConfig, nil
}

func DBConn() (*sql.DB, error) {
	dBConfig, err := LoadDBConfig()
	if err != nil {
		sugar.Errorf("Error %s", err)
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dBConfig.Username, dBConfig.Passwd, dBConfig.Host, dBConfig.Port, dBConfig.DBName)
	sqlDB, err := sql.Open(dBConfig.Driver, dsn)
	if err != nil {
		sugar.Errorf("Error %s", err)
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(1000)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	ct, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	err = sqlDB.PingContext(ct)
	if err != nil {
		sugar.Errorf("Errors %s pinging DB", err)
	}

	//sugar.Infof("==> Connected databse: %s", dsn)

	return sqlDB, nil
}
