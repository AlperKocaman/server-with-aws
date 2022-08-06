package config

import (
	"github.com/spf13/viper"
	"log"
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
	Region string
	S3     S3 `mapstructure:"s3"`
}

type S3 struct {
	BucketName string
}

func InitializeConfig() error {

	log.Println("Initializing config...")

	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into config struct, %v", err)
		return err
	}

	return nil
}

func Get() *Config {
	return config
}