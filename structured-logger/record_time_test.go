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
	"log/slog"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNewRecordTime(t *testing.T) {
	t.Run("format and utc", func(t *testing.T) {
		rt := NewRecordTime(time.RFC3339, true)
		if reflect.ValueOf(rt.valueFunc).Pointer() != reflect.ValueOf(rt.timeUTCFormat).Pointer() {
			t.Errorf("expected %p, got %p", rt.timeUTCFormat, rt.valueFunc)
		}
	})
	t.Run("format", func(t *testing.T) {
		rt := NewRecordTime(time.RFC3339, false)
		if reflect.ValueOf(rt.valueFunc).Pointer() != reflect.ValueOf(rt.timeFormat).Pointer() {
			t.Errorf("expected %p, got %p", rt.timeUTCFormat, rt.valueFunc)
		}
	})
	t.Run("utc", func(t *testing.T) {
		rt := NewRecordTime("", true)
		if reflect.ValueOf(rt.valueFunc).Pointer() != reflect.ValueOf(TimeUTCValue).Pointer() {
			t.Errorf("expected %p, got %p", rt.timeUTCFormat, rt.valueFunc)
		}
	})
	t.Run("none", func(t *testing.T) {
		rt := NewRecordTime("", false)
		if rt.valueFunc != nil {
			t.Errorf("expected %p, got %p", rt.timeUTCFormat, rt.valueFunc)
		}
	})
}

func TestName(t *testing.T) {
	var Logger *slog.Logger

	useUTC, _ := strconv.ParseBool(os.Getenv("LOG_TIME_UTC"))

	recordTime := NewRecordTime(os.Getenv("LOG_TIME_FORMAT"), useUTC)

	options := &slog.HandlerOptions{
		AddSource:   false,
		Level:       GetLevel(os.Getenv("LOG_LEVEL"), slog.LevelWarn),
		ReplaceAttr: recordTime.ReplaceAttr,
	}

	handler := GetHandler(os.Getenv("LOG_HANDLER"), os.Stdout, options, slog.Default().Handler())
	handler = WithProjectAttr("my-project", handler)

	Logger = slog.New(handler)

	Logger.Info("hello", slog.String("user", os.Getenv("USER")))
}
