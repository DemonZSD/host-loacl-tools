package config

import (
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
)
type AppConfig struct {
	PathConfig
}

type PathConfig struct {
	SavePath string
}

var Config *AppConfig
func init()  {
	configPath := "src/config-writer/config/app.ini"
	config, err := ReadConfig(configPath)
	if err != nil{
		return
	}
	Config = config
}

func ReadConfig(configPath string) (appConfig *AppConfig, err error){
	cfg, err := ini.Load(configPath)
	if err != nil {
		return appConfig, errors.New(fmt.Sprintf("parse config file failed: %v", err))
	}
	appConfig = new(AppConfig)
	err = cfg.MapTo(appConfig)
	if err!=nil{
		return appConfig, errors.New(fmt.Sprintf("parse config file failed: %v", err))
	}
	return appConfig, nil
}