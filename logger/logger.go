/*
 * Copyright 2024 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logger

import (
	log_level "github.com/y-du/go-log-level"
	"github.com/y-du/go-log-level/level"
	"log"
	"os"
	"path"
)

type Logger = log_level.Logger

type Level = level.Level

const (
	Off     = level.Off
	Error   = level.Error
	Warning = level.Warning
	Info    = level.Info
	Debug   = level.Debug
	Default = level.Default
)

type LogFileError struct {
	err error
}

func (e *LogFileError) Error() string {
	return e.err.Error()
}

func (e *LogFileError) Unwrap() error {
	return e.err
}

func New(level Level, dirPath, fileName, prefix string, utc, terminal, microseconds bool) (logger *Logger, out *os.File, err error) {
	flags := log.Ldate | log.Ltime | log.Lmsgprefix
	if utc {
		flags = flags | log.LUTC
	}
	if microseconds {
		flags = flags | log.Lmicroseconds
	}
	if terminal {
		out = os.Stderr
	} else {
		out, err = os.OpenFile(path.Join(dirPath, fileName+".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			err = &LogFileError{err: err}
			return
		}
	}
	logger, err = log_level.New(log.New(out, prefix, flags), level)
	return
}

var ParseLevel = level.Parse
