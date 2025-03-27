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
	"log/slog"
	"time"
)

type RecordTime struct {
	format    string
	valueFunc func(t time.Time) slog.Value
}

func NewRecordTime(format string, utc bool) *RecordTime {
	h := RecordTime{format: format}
	switch {
	case format != "" && utc:
		h.valueFunc = h.timeUTCFormat
	case format != "":
		h.valueFunc = h.timeFormat
	case utc:
		h.valueFunc = TimeUTCValue
	}
	return &h
}

func (r *RecordTime) Value(v slog.Value) slog.Value {
	if r.valueFunc == nil {
		return v
	}
	return r.valueFunc(v.Time())
}

func (r *RecordTime) ReplaceAttr(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey && r.valueFunc != nil {
		a.Value = r.valueFunc(a.Value.Time())
	}
	return a
}

func (r *RecordTime) timeFormat(t time.Time) slog.Value {
	return slog.StringValue(t.Format(r.format))
}

func (r *RecordTime) timeUTCFormat(t time.Time) slog.Value {
	return slog.StringValue(t.UTC().Format(r.format))
}

func TimeUTCValue(t time.Time) slog.Value {
	return slog.TimeValue(t.UTC())
}

func TimeUTCReplaceAttr(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey {
		a.Value = slog.TimeValue(a.Value.Time().UTC())
	}
	return a
}
