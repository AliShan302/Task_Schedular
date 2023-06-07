package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration
const (
	DbName = "db.name"
	DbHost = "db.ip"
	DbPort = "db.port"
	DbUser = "db.user"
	DbPass = "db.pass"
)

func init() {
	// env var for db
	_ = viper.BindEnv(DbName, "DB_NAME")
	_ = viper.BindEnv(DbHost, "DB_HOST")
	_ = viper.BindEnv(DbPort, "DB_PORT")
	_ = viper.BindEnv(DbUser, "DB_USER")
	_ = viper.BindEnv(DbPass, "DB_PASS")

	// defaults
	viper.SetDefault(DbName, "task_schedular")
	viper.SetDefault(DbHost, "localhost")
	viper.SetDefault(DbPort, "3306")

}
