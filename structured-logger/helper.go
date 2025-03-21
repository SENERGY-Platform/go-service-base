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

package structured_logger

import (
	"io"
	"log/slog"
	"os"
)

const (
	levelDebug = "debug"
	levelInfo  = "info"
	levelWarn  = "warn"
	levelError = "error"
)

func GetLevel(value string, defaultLevel slog.Leveler) slog.Leveler {
	switch value {
	case levelDebug:
		return slog.LevelDebug
	case levelInfo:
		return slog.LevelInfo
	case levelWarn:
		return slog.LevelWarn
	case levelError:
		return slog.LevelError
	default:
		return defaultLevel
	}
}

const (
	textHandlerSelector    = "text"
	jsonHandlerSelector    = "json"
	discardHandlerSelector = "discard"
)

func GetHandler(value string, writer io.Writer, opts *slog.HandlerOptions, defaultHandler slog.Handler) slog.Handler {
	switch value {
	case textHandlerSelector:
		return slog.NewTextHandler(writer, opts)
	case jsonHandlerSelector:
		return slog.NewJSONHandler(writer, opts)
	case discardHandlerSelector:
		return slog.DiscardHandler
	default:
		return defaultHandler
	}
}

func GetLogFile(filePath string, filePerm os.FileMode) (io.WriteCloser, error) {
	return os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePerm)
}

const projectKey = "project"

func WithProjectAttr(name string, handler slog.Handler) slog.Handler {
	if name != "" {
		return handler.WithAttrs([]slog.Attr{{
			Key:   projectKey,
			Value: slog.StringValue(name),
		}})
	}
	return handler
}
