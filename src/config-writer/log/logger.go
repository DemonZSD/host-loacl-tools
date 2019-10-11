// Copyright 2019 Weshzhu
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// author: Weshzhu
package logger

import (
	"config-writer/config"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
	"log"
	"fmt"
	"github.com/pkg/errors"
)

var logger *logrus.Logger

func init() {
	logr, err := initLogger()
	if err != nil {
		log.Println(fmt.Sprintf("init logger failed: %v", err))
	}
	logger = logr
}

func initLogger() (*logrus.Logger, error) {
	var cfg = config.Appcfg
	if cfg == nil {
		return nil, errors.New("get config error")
	}
	var logDir = cfg.LogPath
	if _, err := os.Stat(logDir); err != nil {
		err := os.Mkdir(logDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	writer, err := rotatelogs.New(
		logDir+"/%Y%m%dhost-local-tools.log",
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithMaxAge(30*24*time.Hour),
	)
	if err != nil {
		return nil, err
	}
	logging := logrus.New()
	logging.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: writer,
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.PanicLevel: writer,
		},
		&logrus.TextFormatter{},
	))
	return logging, nil
}

func GetLog() *logrus.Logger{
	return logger
}