package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func LoadConfiguration(configFile string) (*Configuration, error) {
	vpr := viper.New()
	vpr.SetConfigFile(configFile)
	if err := vpr.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Configuration
	err := vpr.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	validate := validator.New()
	if err = validate.Struct(&config); err != nil {
		return nil, err
	}

	vpr.WatchConfig()
	vpr.OnConfigChange(func(e fsnotify.Event) {
		if err := vpr.Unmarshal(&config); err != nil {
			log.Print("failed to update public config after hot reload", "err", err)
		}
		//loadSecret(&config)
	})

	return &config, nil
}
