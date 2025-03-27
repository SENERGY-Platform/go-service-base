/*
 * Copyright 2025 InfAI (CC SES)
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

package struct_logger

import (
	"github.com/SENERGY-Platform/go-service-base/struct-logger/attributes"
	"io"
	"log/slog"
	"os"
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
)

func GetLevel(value string, defaultLevel slog.Leveler) slog.Leveler {
	switch value {
	case LevelDebug:
		return slog.LevelDebug
	case LevelInfo:
		return slog.LevelInfo
	case LevelWarn:
		return slog.LevelWarn
	case LevelError:
		return slog.LevelError
	default:
		return defaultLevel
	}
}

const (
	TextHandlerSelector    = "text"
	JsonHandlerSelector    = "json"
	DiscardHandlerSelector = "discard"
)

func GetHandler(value string, writer io.Writer, opts *slog.HandlerOptions, defaultHandler slog.Handler) slog.Handler {
	switch value {
	case TextHandlerSelector:
		return slog.NewTextHandler(writer, opts)
	case JsonHandlerSelector:
		return slog.NewJSONHandler(writer, opts)
	case DiscardHandlerSelector:
		return slog.DiscardHandler
	default:
		return defaultHandler
	}
}

func GetLogFile(filePath string, filePerm os.FileMode) (io.WriteCloser, error) {
	return os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePerm)
}

type Config struct {
	Handler    string `json:"handler" env_var:"LOGGER_HANDLER"`
	Level      string `json:"level" env_var:"LOGGER_LEVEL"`
	TimeFormat string `json:"time_format" env_var:"LOGGER_TIME_FORMAT"`
	TimeUtc    bool   `json:"time_utc" env_var:"LOGGER_TIME_UTC"`
	FilePath   string `json:"file_path" env_var:"LOGGER_FILE_PATH"`
	AddSource  bool   `json:"add_source" env_var:"LOGGER_ADD_SOURCE"`
	AddMeta    bool   `json:"add_meta" env_var:"LOGGER_ADD_META"`
}

func New(c Config, out io.Writer, organization, project string) *slog.Logger {
	recordTime := NewRecordTime(c.TimeFormat, c.TimeUtc)
	options := &slog.HandlerOptions{
		AddSource:   c.AddSource,
		Level:       GetLevel(c.Level, slog.LevelInfo),
		ReplaceAttr: recordTime.ReplaceAttr,
	}
	handler := GetHandler(c.Handler, out, options, slog.Default().Handler())
	if c.AddMeta {
		var attr []slog.Attr
		if organization != "" {
			attr = append(attr, slog.String(attributes.OrganizationKey, organization))
		}
		if project != "" {
			attr = append(attr, slog.String(attributes.ProjectKey, project))
		}
		handler = handler.WithAttrs(attr)
	}
	return slog.New(handler)
}
