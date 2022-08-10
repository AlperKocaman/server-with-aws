package config

import (
	"github.com/spf13/viper"
	"log"
)

const (
	AppConfig  = "config"
	TestConfig = "config-test"
)

var config *Config

type Config struct {
	Server Server `mapstructure:"server"`
	AWS    AWS    `mapstructure:"aws"`
}

type Server struct {
	Host string
	Port string
}

type AWS struct {
	Region         string
	AccessKeyID    string
	SecretAcessKey string
	SessionToken   string
	S3             S3 `mapstructure:"s3"`
}

type S3 struct {
	Bucket string
}

func InitializeConfigForApp() error {
	return InitializeConfig(AppConfig)
}

func InitializeConfigForTest() error {
	return InitializeConfig(TestConfig)
}

func InitializeConfig(configName string) error {

	log.Println("Initializing config...")

	viper.SetConfigName(configName)
	viper.AddConfigPath("./cmd/config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
		return err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Printf("Unable to decode into config struct, %v", err)
		return err
	}

	return nil
}

func Get() *Config {
	return config
}
