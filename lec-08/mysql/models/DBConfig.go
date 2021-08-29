package models

import "fmt"

type DBConfig struct {
	Driver   string `mapstructure:"DB_DRIVER"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	Username string `mapstructure:"DB_USERNAME"`
	Passwd   string `mapstructure:"DB_PASSWORD"`
	DBName   string `mapstructure:"DB_NAME"`
}

func (dBConfig DBConfig) String() string {
	return fmt.Sprintf("[Driver: %s, Host: %s, Port: %s, Username: %s, Passwd: %s, DBName: %s]",
		dBConfig.Driver, dBConfig.Host, dBConfig.Port, dBConfig.Username, dBConfig.Passwd, dBConfig.DBName)
}
