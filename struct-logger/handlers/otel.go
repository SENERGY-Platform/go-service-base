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

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/trace"
)

type OtelHandler struct {
	handler slog.Handler
}

func (h *OtelHandler) Handle(ctx context.Context, r slog.Record) error {
	for attr := range ctxValFunc(ctx) {
		r.AddAttrs(attr)
	}
	gc, ok := ctx.(*gin.Context)
	if ok {
		ctx = gc.Request.Context()
	}

	span := trace.SpanFromContext(ctx)
	if span.IsRecording() {
		attrs := []attribute.KeyValue{}
		r.Attrs(func(attr slog.Attr) bool {
			attrs = append(attrs, attribute.KeyValue{
				Key:   attribute.Key(attr.Key),
				Value: attribute.StringValue(attr.Value.String()),
			})
			return true
		})
		span.AddEvent(r.Message, trace.WithAttributes(attrs...))
	}
	return h.handler.Handle(ctx, r)
}

func (h *OtelHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *OtelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &OtelHandler{
		handler: h.handler.WithAttrs(attrs),
	}
}

func (h *OtelHandler) WithGroup(name string) slog.Handler {
	return &OtelHandler{
		handler: h.handler.WithGroup(name),
	}
}

func NewOpenTelemetryHandler(baseHandler slog.Handler) slog.Handler {
	return &OtelHandler{
		handler: NewContextHandler(baseHandler, ctxValFunc)}
}

func ctxValFunc(ctx context.Context) iter.Seq[slog.Attr] {
	return func(yield func(slog.Attr) bool) {
		gc, ok := ctx.(*gin.Context)
		if ok {
			ctx = gc.Request.Context()
		}

		bag := baggage.FromContext(ctx)
		for _, m := range bag.Members() {
			if !yield(slog.String(m.Key(), m.Value())) {
				return
			}
		}
	}
}
