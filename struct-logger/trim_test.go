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
	"bytes"
	"reflect"
	"strings"
	"testing"
	"time"
)
import "testing/synctest"

func TestTrim(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{in: "abcdefghijklmnopqrstuvwxyz", out: "abc[...]yz"},
		{in: "abcdefghijk", out: "abc[...]jk"},
		{in: "abcdefghij", out: "abcdefghij"},
		{in: "abcdefghi", out: "abcdefghi"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Trim(tt.in, "[...]", 3, 2); got != tt.out {
				t.Errorf("Trim() = %v, want %v", got, tt.out)
			}
		})
	}
}

func TestLoggerWithTrim(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		buff := bytes.NewBufferString("")
		log := New(Config{
			Handler:        JsonHandlerSelector,
			Level:          "debug",
			TimeFormat:     time.RFC3339Nano,
			TimeUtc:        true,
			TrimFormat:     "3:[...]:2",
			TrimAttributes: "foo,bar",
			AddMeta:        true,
		}, buff, "org", "trim")
		withoutTrim := New(Config{
			Handler:    JsonHandlerSelector,
			Level:      "debug",
			TimeFormat: time.RFC3339Nano,
			TimeUtc:    true,
			AddMeta:    true,
		}, buff, "org", "notrim")

		for _, msg := range []string{
			"abcdefghijklmnopqrstuvwxyz",
			"abcdefghijk",
			"abcdefghij",
			"abcdefghi",
		} {
			log.Error(msg, "foo", msg, "bar", msg, "batz", msg, "num", 13, "obj", map[string]interface{}{"foo": true, "bar": 42})
			withoutTrim.Error(msg, "foo", msg, "bar", msg, "batz", msg, "num", 13, "obj", map[string]interface{}{"foo": true, "bar": 42})
		}

		expected := []string{
			`{"time":"2000-01-01T00:00:00Z","level":"ERROR","msg":"abc[...]yz","organization":"org","project":"trim","foo":"abc[...]yz","bar":"abc[...]yz","batz":"abcdefghijklmnopqrstuvwxyz","num":13,"obj":{"bar":42,"foo":true}}`,
			`{"time":"2000-01-01T00:00:00Z","level":"ERROR","msg":"abcdefghijklmnopqrstuvwxyz","organization":"org","project":"notrim","foo":"abcdefghijklmnopqrstuvwxyz","bar":"abcdefghijklmnopqrstuvwxyz","batz":"abcdefghijklmnopqrstuvwxyz","num":13,"obj":{"bar":42,"foo":true}}`,
			`{"time":"2000-01-01T00:00:00Z","level":"ERROR","msg":"abc[...]jk","organization":"org","project":"trim","foo":"abc[...]jk","bar":"abc[...]jk","batz":"abcdefghijk","num":13,"obj":{"bar":42,"foo":true}}`,
			`{"time":"2000-01-01T00:00:00Z","level":"ERROR","msg":"abcdefghijk","organization":"org","project":"notrim","foo":"abcdefghijk","bar":"abcdefghijk","batz":"abcdefghijk","num":13,"obj":{"bar":42,"foo":true}}`,
			`{"time":"2000-01-01T00:00:00Z","level":"ERROR","msg":"abcdefghij","organization":"org","project":"trim","foo":"abcdefghij","bar":"abcdefghij","batz":"abcdefghij","num":13,"obj":{"bar":42,"foo":true}}`,
			`{"time":"2000-01-01T00:00:00Z","level":"ERROR","msg":"abcdefghij","organization":"org","project":"notrim","foo":"abcdefghij","bar":"abcdefghij","batz":"abcdefghij","num":13,"obj":{"bar":42,"foo":true}}`,
			`{"time":"2000-01-01T00:00:00Z","level":"ERROR","msg":"abcdefghi","organization":"org","project":"trim","foo":"abcdefghi","bar":"abcdefghi","batz":"abcdefghi","num":13,"obj":{"bar":42,"foo":true}}`,
			`{"time":"2000-01-01T00:00:00Z","level":"ERROR","msg":"abcdefghi","organization":"org","project":"notrim","foo":"abcdefghi","bar":"abcdefghi","batz":"abcdefghi","num":13,"obj":{"bar":42,"foo":true}}`,
		}
		actual := []string{}
		for out := range strings.Lines(buff.String()) {
			actual = append(actual, strings.TrimSpace(out))
		}
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("\ne %v\na %v", expected, actual)
		}
	})
}
