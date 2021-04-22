package config

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"sync"
)

type MongoDbConfig struct {
	Driver   string `mapstructure:"Driver"`
	Name     string `mapstructure:"Name"`
	Host     string `mapstructure:"Host"`
	Port     int    `mapstructure:"Port"`
	Username string `mapstructure:"Username"`
	Password string `mapstructure:"Password"`
}

type AppConfig struct {
	Name          string        `mapstructure:"Name"`
	Port          int           `mapstructure:"Port"`
	SecretKey     string        `mapstructure:"SecretKey"`
	MasterMongoDb MongoDbConfig `mapstructure:"MasterMongoDb"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfigs() *AppConfig {
	if appConfig != nil {
		return appConfig
	}

	lock.Lock()
	defer lock.Unlock()

	if appConfig != nil {
		return appConfig
	}

	appConfig = initConfig()

	return appConfig
}

func initConfig() *AppConfig {
	var finalConfig AppConfig
	viper.BindEnv("Name", "APP_NAME")
	viper.BindEnv("Port", "APP_PORT")
	viper.BindEnv("SecretKey", "APP_SECRET_KEY")
	viper.BindEnv("MasterMongoDb.Driver", "MASTER_MONGO_DB_DRIVER")
	viper.BindEnv("MasterMongoDb.Name", "MASTER_MONGO_DB_NAME")
	viper.BindEnv("MasterMongoDb.Host", "MASTER_MONGO_DB_HOST")
	viper.BindEnv("MasterMongoDb.Port", "MASTER_MONGO_DB_PORT")
	viper.BindEnv("MasterMongoDb.Username", "MASTER_MONGO_DB_USERNAME")
	viper.BindEnv("MasterMongoDb.Password", "MASTER_MONGO_DB_PASSWORD")
	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract config, will use default value")
	}
	return &finalConfig
}
