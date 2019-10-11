package config

import (
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
)

type AppConfig struct {
	PathConfig
	VFConfig
	IPAllocate
	EtcdConfig
	LogConfig
}

type PathConfig struct {
	SavePath string
}
type LogConfig struct {
	LogPath string
}
type VFConfig struct {
	VFName string
	Name string
	Type string
	Mode string
}
type IPAllocate struct {
	Subnet string
}
type EtcdConfig struct {
	EtcdAddr string
}

var Appcfg *AppConfig

func init() {
	configPath := "resource/app.ini"
	config, err := ReadConfig(configPath)
	if err != nil {
		return
	}
	Appcfg = config
}

func ReadConfig(configPath string) (appConfig *AppConfig, err error) {
	cfg, err := ini.Load(configPath)
	if err != nil {
		return appConfig, errors.New(fmt.Sprintf("load config file failed: %v", err))
	}
	appConfig = new(AppConfig)
	err = cfg.MapTo(appConfig)
	if err != nil {
		return appConfig, errors.New(fmt.Sprintf("parse config file failed: %v", err))
	}
	return appConfig, nil
}
