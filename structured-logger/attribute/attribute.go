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
	MethodAttrKey        = "method"
	StatusCodeAttrKey    = "status"
	LatencyAttrKey       = "latency"
	PathAttrKey          = "path"
	ProtocolAttrKey      = "protocol"
	UserAgentAttrKey     = "user_agent"
	BodySizeAttrKey      = "body_size"
	HeadersAttrKey       = "headers"
	BodyAttrKey          = "body"
	ErrorAttrKey         = "error"
	StackTraceAttrKey    = "stack_trace"
	ProjectAttrKey       = "project"
	OrganizationAttrKey  = "organization"
	LogRecordTypeAttrKey = "log_record_type"
)

const (
	HttpAccessLogRecordType = "http_access"
)

var Provider = provider{}

type provider struct{}

func (p *provider) PathKey() string {
	return PathAttrKey
}

func (p *provider) StatusCodeKey() string {
	return StatusCodeAttrKey
}

func (p *provider) MethodKey() string {
	return MethodAttrKey
}

func (p *provider) LatencyKey() string {
	return LatencyAttrKey
}

func (p *provider) ProtocolKey() string {
	return ProtocolAttrKey
}

func (p *provider) UserAgentKey() string {
	return UserAgentAttrKey
}

func (p *provider) BodySizeKey() string {
	return BodySizeAttrKey
}

func (p *provider) HeadersKey() string {
	return HeadersAttrKey
}

func (p *provider) BodyKey() string {
	return BodyAttrKey
}

func (p *provider) ErrorKey() string {
	return ErrorAttrKey
}

func (p *provider) StackTraceKey() string {
	return StackTraceAttrKey
}

func (p *provider) ProjectKey() string {
	return ProjectAttrKey
}

func (p *provider) OrganizationKey() string {
	return OrganizationAttrKey
}

func (p *provider) LogRecordTypeKey() string {
	return LogRecordTypeAttrKey
}

func ErrorAttr(err error) slog.Attr {
	return slog.String(ErrorAttrKey, err.Error())
}

func ProjectAttr(val string) slog.Attr {
	return slog.String(ProjectAttrKey, val)
}

func OrganizationAttr(val string) slog.Attr {
	return slog.String(OrganizationAttrKey, val)
}

func HttpAccessLogRecordTypeAttr() slog.Attr {
	return slog.String(LogRecordTypeAttrKey, HttpAccessLogRecordType)
}
