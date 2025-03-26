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

package attribute

import "log/slog"

const (
	MethodAttrKey       = "method"
	StatusCodeAttrKey   = "status"
	LatencyAttrKey      = "latency"
	PathAttrKey         = "path"
	ProtocolAttrKey     = "protocol"
	UserAgentAttrKey    = "user_agent"
	BodySizeAttrKey     = "body_size"
	HeadersAttrKey      = "headers"
	BodyAttrKey         = "body"
	ErrorMessageAttrKey = "error"
	StackTraceAttrKey   = "stack_trace"
	ProjectAttrKey      = "project"
	OrganizationAttrKey = "organization"
	LogRecordTypeKey    = "log_record_type"
)

const (
	HttpAccessLogRecordType = "http_access"
)

var Provider = provider{}

type provider struct{}

func (p *provider) AppendPath(args []any, val any) []any {
	return append(args, slog.Attr{Key: PathAttrKey, Value: slog.AnyValue(val)})
}

func (p *provider) PathKey() string {
	return PathAttrKey
}

func (p *provider) AppendStatusCode(args []any, val any) []any {
	return append(args, slog.Attr{Key: StatusCodeAttrKey, Value: slog.AnyValue(val)})
}

func (p *provider) StatusCodeKey() string {
	return StatusCodeAttrKey
}

func (p *provider) AppendMethod(args []any, val any) []any {
	return append(args, slog.Attr{Key: MethodAttrKey, Value: slog.AnyValue(val)})
}

func (p *provider) MethodKey() string {
	return MethodAttrKey
}

func (p *provider) AppendLatency(args []any, val any) []any {
	return append(args, slog.Attr{Key: LatencyAttrKey, Value: slog.AnyValue(val)})
}

func (p *provider) LatencyKey() string {
	return LatencyAttrKey
}

func (p *provider) AppendProtocol(args []any, val any) []any {
	return append(args, slog.Attr{Key: ProtocolAttrKey, Value: slog.AnyValue(val)})
}

func (p *provider) ProtocolKey() string {
	return ProtocolAttrKey
}

func (p *provider) AppendUserAgent(args []any, val any) []any {
	return append(args, slog.Attr{Key: UserAgentAttrKey, Value: slog.AnyValue(val)})
}

func (p *provider) UserAgentKey() string {
	return UserAgentAttrKey
}

func (p *provider) AppendBodySize(args []any, val any) []any {
	return append(args, slog.Attr{Key: BodySizeAttrKey, Value: slog.AnyValue(val)})
}

func (p *provider) BodySizeKey() string {
	return BodySizeAttrKey
}

func (p *provider) AppendHeaders(args []any, val any) []any {
	return append(args, slog.Attr{Key: HeadersAttrKey, Value: slog.AnyValue(val)})
}

func (p *provider) HeadersKey() string {
	return HeadersAttrKey
}

func (p *provider) AppendBody(args []any, val any) []any {
	return append(args, slog.Attr{Key: BodyAttrKey, Value: slog.AnyValue(val)})
}

func (p *provider) BodyKey() string {
	return BodyAttrKey
}

func (p *provider) AppendErrMsg(args []any, val any) []any {
	return append(args, slog.Attr{Key: ErrorMessageAttrKey, Value: slog.AnyValue(val)})
}

func (p *provider) ErrMsgKey() string {
	return ErrorMessageAttrKey
}

func (p *provider) AppendStackTrace(args []any, val any) []any {
	return append(args, slog.Attr{Key: StackTraceAttrKey, Value: slog.AnyValue(val)})
}

func (p *provider) StackTraceKey() string {
	return StackTraceAttrKey
}

func (p *provider) AppendProject(args []any, val any) []any {
	return append(args, slog.Attr{Key: ProjectAttrKey, Value: slog.AnyValue(val)})
}

func (p *provider) AppendProjectAttr(args []slog.Attr, val any) []slog.Attr {
	return append(args, slog.Attr{Key: ProjectAttrKey, Value: slog.AnyValue(val)})
}

func (p *provider) ProjectAttr(val any) slog.Attr {
	return slog.Attr{Key: ProjectAttrKey, Value: slog.AnyValue(val)}
}

func (p *provider) AppendOrganization(args []any, val any) []any {
	return append(args, slog.Attr{Key: OrganizationAttrKey, Value: slog.AnyValue(val)})
}

func (p *provider) AppendOrganizationAttr(args []slog.Attr, val any) []slog.Attr {
	return append(args, slog.Attr{Key: OrganizationAttrKey, Value: slog.AnyValue(val)})
}

func (p *provider) OrganizationAttr(val any) slog.Attr {
	return slog.Attr{Key: OrganizationAttrKey, Value: slog.AnyValue(val)}
}

func (p *provider) AppendHttpAccessLogRecordType(args []any) []any {
	return append(args, slog.String(LogRecordTypeKey, HttpAccessLogRecordType))
}

func (p *provider) AppendHttpAccessLogRecordTypeAttr(args []slog.Attr) []slog.Attr {
	return append(args, slog.String(LogRecordTypeKey, HttpAccessLogRecordType))
}

func (p *provider) HttpAccessLogRecordTypeAttr() slog.Attr {
	return slog.String(LogRecordTypeKey, HttpAccessLogRecordType)
}

func (p *provider) HttpAccessLogRecordTypeKeyVal() (string, string) {
	return LogRecordTypeKey, HttpAccessLogRecordType
}
