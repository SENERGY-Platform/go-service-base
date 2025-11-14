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

package handlers

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"sync"
)

var levelColor = map[slog.Level]string{
	slog.LevelDebug: "\033[36m", // cyan
	slog.LevelInfo:  "\033[32m", // green
	slog.LevelWarn:  "\033[33m", // yellow
	slog.LevelError: "\033[31m", // red
}

const reset = "\033[0m"

// colorHandler wraps another slog.Handler and writes ANSI color codes
// to the writer before and after the wrapped handler emits the log line.
// Use the same io.Writer that the wrapped handler writes to (e.g. os.Stdout).
type colorHandler struct {
	h  slog.Handler
	w  io.Writer
	mu sync.Mutex
}

func NewColorHandler(w io.Writer, wrapped slog.Handler) slog.Handler {
	return &colorHandler{h: wrapped, w: w}
}

func (c *colorHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return c.h.Enabled(ctx, level)
}

func (c *colorHandler) Handle(ctx context.Context, r slog.Record) error {
	// Choose color based on level; default to no color.
	color, ok := levelColor[r.Level]
	if !ok {
		color = ""
	}

	// Lock to avoid interleaving color/reset between concurrent writes.
	c.mu.Lock()
	defer c.mu.Unlock()

	// Write color prefix (if any).
	if color != "" {
		if _, err := fmt.Fprint(c.w, color); err != nil {
			// ignore write error and continue to let inner handler run
		}
	}

	// Let the wrapped handler write the log line.
	err := c.h.Handle(ctx, r)

	// Write reset suffix (if color was used).
	if color != "" {
		if _, werr := fmt.Fprint(c.w, reset); werr != nil && err == nil {
			// if writing reset failed but handler reported no error, surface the write error
			err = werr
		}
	}

	return err
}

func (c *colorHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &colorHandler{
		h: c.h.WithAttrs(attrs),
		w: c.w,
	}
}

func (c *colorHandler) WithGroup(name string) slog.Handler {
	return &colorHandler{
		h: c.h.WithGroup(name),
		w: c.w,
	}
}
