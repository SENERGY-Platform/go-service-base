/*
 * Copyright 2026 InfAI (CC SES)
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
	"iter"
	"log/slog"
)

type ContextValueFunc func(ctx context.Context) iter.Seq[slog.Attr]

type ContextHandler struct {
	slog.Handler
	ctxValF ContextValueFunc
}

func NewContextHandler(baseHandler slog.Handler, ctxValFunc ContextValueFunc) slog.Handler {
	return &ContextHandler{
		Handler: baseHandler,
		ctxValF: ctxValFunc,
	}
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	for attr := range h.ctxValF(ctx) {
		r.AddAttrs(attr)
	}
	return h.Handler.Handle(ctx, r)
}

func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextHandler{
		Handler: h.Handler.WithAttrs(attrs),
		ctxValF: h.ctxValF,
	}
}

func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return &ContextHandler{
		Handler: h.Handler.WithGroup(name),
		ctxValF: h.ctxValF,
	}
}
