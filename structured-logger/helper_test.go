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
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path"
	"reflect"
	"testing"
)

func TestGetLevel(t *testing.T) {
	t.Run("debug", func(t *testing.T) {
		lvl := GetLevel(LevelDebug, slog.LevelWarn)
		if lvl != slog.LevelDebug {
			t.Errorf("expected %s, got %s", slog.LevelDebug, lvl)
		}
	})
	t.Run("info", func(t *testing.T) {
		lvl := GetLevel(LevelInfo, slog.LevelWarn)
		if lvl != slog.LevelInfo {
			t.Errorf("expected %s, got %s", slog.LevelInfo, lvl)
		}
	})
	t.Run("warn", func(t *testing.T) {
		lvl := GetLevel(LevelWarn, slog.LevelWarn)
		if lvl != slog.LevelWarn {
			t.Errorf("expected %s, got %s", slog.LevelWarn, lvl)
		}
	})
	t.Run("error", func(t *testing.T) {
		lvl := GetLevel(LevelError, slog.LevelWarn)
		if lvl != slog.LevelError {
			t.Errorf("expected %s, got %s", slog.LevelError, lvl)
		}
	})
	t.Run("default", func(t *testing.T) {
		lvl := GetLevel("", slog.LevelWarn)
		if lvl != slog.LevelWarn {
			t.Errorf("expected %s, got %s", slog.LevelWarn, lvl)
		}
	})
}

func TestGetHandler(t *testing.T) {
	t.Run("text", func(t *testing.T) {
		h := GetHandler(TextHandlerSelector, nil, nil, slog.Default().Handler())
		if reflect.TypeOf(h) != reflect.TypeOf(&slog.TextHandler{}) {
			t.Errorf("expected %s, got %s", reflect.TypeOf(&slog.TextHandler{}), reflect.TypeOf(h))
		}
	})
	t.Run("json", func(t *testing.T) {
		h := GetHandler(JsonHandlerSelector, nil, nil, slog.Default().Handler())
		if reflect.TypeOf(h) != reflect.TypeOf(&slog.JSONHandler{}) {
			t.Errorf("expected %s, got %s", reflect.TypeOf(&slog.JSONHandler{}), reflect.TypeOf(h))
		}
	})
	t.Run("discard", func(t *testing.T) {
		h := GetHandler(DiscardHandlerSelector, nil, nil, slog.Default().Handler())
		if reflect.TypeOf(h) != reflect.TypeOf(slog.DiscardHandler) {
			t.Errorf("expected %s, got %s", reflect.TypeOf(slog.DiscardHandler), reflect.TypeOf(h))
		}
	})
	t.Run("default", func(t *testing.T) {
		h := GetHandler("", nil, nil, slog.Default().Handler())
		if !reflect.DeepEqual(h, slog.Default().Handler()) {
			t.Errorf("expected %v, got %v", slog.Default().Handler(), h)
		}
	})
}

func TestGetLogFile(t *testing.T) {
	tmpDir := t.TempDir()
	t.Run("create and write", func(t *testing.T) {
		file, err := GetLogFile(path.Join(tmpDir, "test.log"), 0644)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()
		a := []string{
			"line 0",
			"line 1",
		}
		for _, s := range a {
			_, err = fmt.Fprintln(file, s)
			if err != nil {
				t.Fatal(err)
			}
		}
		if err = file.Close(); err != nil {
			t.Fatal(err)
		}
		file2, err := os.Open(path.Join(tmpDir, "test.log"))
		if err != nil {
			t.Fatal(err)
		}
		defer file2.Close()
		scanner := bufio.NewScanner(file2)
		var b []string
		for scanner.Scan() {
			b = append(b, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		if !reflect.DeepEqual(a, b) {
			t.Errorf("expected %v, got %v", a, b)
		}
	})
	t.Run("open and write", func(t *testing.T) {
		file, err := GetLogFile(path.Join(tmpDir, "test.log"), 0644)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()
		_, err = fmt.Fprintln(file, "line 2")
		if err != nil {
			t.Fatal(err)
		}
		a := []string{
			"line 0",
			"line 1",
			"line 2",
		}
		if err = file.Close(); err != nil {
			t.Fatal(err)
		}
		file2, err := os.Open(path.Join(tmpDir, "test.log"))
		if err != nil {
			t.Fatal(err)
		}
		defer file2.Close()
		scanner := bufio.NewScanner(file2)
		var b []string
		for scanner.Scan() {
			b = append(b, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		if !reflect.DeepEqual(a, b) {
			t.Errorf("expected %v, got %v", a, b)
		}
	})
}
