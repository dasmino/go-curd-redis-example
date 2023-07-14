package config

import "github.com/spf13/viper"

type Config struct {
	// MySQL Setup
	DBHost     string `mapstructure:"SQL_HOST"`
	DBUsername string `maptructure:"SQL_USER"`
	DBPassword string `maptructure:"SQL_PASSWORD"`
	DBName     string `maptructure:"SQL_DB"`
	DBPort     string `maptructure:"SQL_PORT"`

	// Redis Setup
	RedisUrl string `maptructure:"REDIS_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	// handle null
	viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
