package configuration

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DATABASE_URL"`
	DBName     string `mapstructure:"DATABASE_DB"`
	DBUser     string `mapstructure:"DATABASE_USER"`
	DBPort     string `mapstructure:"DATABASE_PORT"`
	DBPassword string `mapstructure:"DATABASE_PASSWORD"`
}

var envs = []string{
	"DATABASE_URL", "DATABASE_DB", "DATABASE_USER", "DATABASE_PORT", "DATABASE_PASSWORD",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("../");
	viper.SetConfigType("env");
	viper.SetConfigName("app");
	
	err := viper.ReadInConfig()

	if err != nil {
		return config, err;
	}

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}