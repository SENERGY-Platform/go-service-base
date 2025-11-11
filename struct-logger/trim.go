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
	"context"
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"strings"
)

func Trim(s string, ellipse string, startLimit int, endLimit int) string {
	if len(s) <= startLimit+endLimit+len(ellipse) {
		return s
	}
	return s[0:startLimit] + ellipse + s[len(s)-endLimit:]
}

func ApplyTrimFormat(handler slog.Handler, format string, attributes []string) (result slog.Handler, err error) {
	if format == "" {
		return handler, nil
	}
	parts := strings.Split(format, ":")
	if len(parts) < 3 {
		return handler, errors.New("invalid trim format")
	}
	trimHandler := &TrimHandler{
		Parent: handler,
	}
	trimHandler.StartLimit, err = strconv.Atoi(parts[0])
	if err != nil {
		return handler, fmt.Errorf("invalid trim format: %w", err)
	}
	trimHandler.EndLimit, err = strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		return handler, fmt.Errorf("invalid trim format: %w", err)
	}
	trimHandler.Ellipse = strings.Join(parts[1:len(parts)-1], ":")
	trimHandler.TrimmedAttributes = attributes
	return trimHandler, nil
}

type TrimHandler struct {
	StartLimit        int
	EndLimit          int
	Ellipse           string
	TrimmedAttributes []string
	Parent            slog.Handler
}

func (this *TrimHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return this.Parent.Enabled(ctx, level)
}

func (this *TrimHandler) Handle(ctx context.Context, record slog.Record) error {
	updated := slog.NewRecord(record.Time, record.Level, Trim(record.Message, this.Ellipse, this.StartLimit, this.EndLimit), record.PC)
	record.Attrs(func(attr slog.Attr) bool {
		if slices.Contains(this.TrimmedAttributes, attr.Key) {
			updated.AddAttrs(slog.String(attr.Key, Trim(attr.Value.String(), this.Ellipse, this.StartLimit, this.EndLimit)))
		} else {
			updated.AddAttrs(attr)
		}
		return true
	})
	return this.Parent.Handle(ctx, updated)
}

func (this *TrimHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return this.Parent.WithAttrs(attrs)
}

func (this *TrimHandler) WithGroup(name string) slog.Handler {
	return this.Parent.WithGroup(name)
}
