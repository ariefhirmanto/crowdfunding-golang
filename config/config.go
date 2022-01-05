package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type MainConfig struct {
	Database DatabaseConfigurations
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	Host       string `mapstructure:"Host"`
	Port       string `mapstructure:"Port"`
	DBName     string `mapstructure:"DBName"`
	DBUser     string `mapstructure:"DBUser"`
	DBPassword string `mapstructure:"DBPassword"`
}

func LoadConfig(path string) (config MainConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	if err := viper.BindEnv("Database"); err != nil {
		return config, err
	}

	err = viper.Unmarshal(&config)
	fmt.Println(err)
	fmt.Printf("%+v\n", config)
	if err != nil {
		fmt.Printf("Failed to unmarshal")
		return config, err
	}

	return config, nil
}
