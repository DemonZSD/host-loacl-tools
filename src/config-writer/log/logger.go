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
package log

import (
	"config-writer/config"
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Logger *logrus.Logger

func init() {
	log, err := initLogger()
	if err != nil {
		log.Errorln(fmt.Sprintf("init logger failed: %v", err))
	}
	Logger = log
}

func initLogger() (*logrus.Logger, error) {
	var logDir = config.Appcfg.LogPath
	if _, err := os.Stat(logDir); err != nil {
		err := os.Mkdir(logDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	writer, err := rotatelogs.New(
		logDir+"host-local-tools-log%Y%m%d",
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithMaxAge(30*24*time.Hour),
	)
	if err != nil {
		return nil, err
	}
	log := logrus.New()
	log.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: writer,
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.PanicLevel: writer,
		},
		&logrus.TextFormatter{},
	))
	return log, nil
}
