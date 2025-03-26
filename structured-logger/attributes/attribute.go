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

package attributes

import "log/slog"

const (
	MethodKey        = "method"
	StatusCodeKey    = "status"
	LatencyKey       = "latency"
	PathKey          = "path"
	ProtocolKey      = "protocol"
	UserAgentKey     = "user_agent"
	BodySizeKey      = "body_size"
	HeadersKey       = "headers"
	BodyKey          = "body"
	ErrorKey         = "error"
	StackTraceKey    = "stack_trace"
	ProjectKey       = "project"
	OrganizationKey  = "organization"
	LogRecordTypeKey = "log_record_type"
)

const (
	HttpAccessLogRecordTypeVal = "http_access"
)

var Provider = provider{}

type provider struct{}

func (p *provider) PathKey() string {
	return PathKey
}

func (p *provider) StatusCodeKey() string {
	return StatusCodeKey
}

func (p *provider) MethodKey() string {
	return MethodKey
}

func (p *provider) LatencyKey() string {
	return LatencyKey
}

func (p *provider) ProtocolKey() string {
	return ProtocolKey
}

func (p *provider) UserAgentKey() string {
	return UserAgentKey
}

func (p *provider) BodySizeKey() string {
	return BodySizeKey
}

func (p *provider) HeadersKey() string {
	return HeadersKey
}

func (p *provider) BodyKey() string {
	return BodyKey
}

func (p *provider) ErrorKey() string {
	return ErrorKey
}

func (p *provider) StackTraceKey() string {
	return StackTraceKey
}

func (p *provider) ProjectKey() string {
	return ProjectKey
}

func (p *provider) OrganizationKey() string {
	return OrganizationKey
}

func (p *provider) LogRecordTypeKey() string {
	return LogRecordTypeKey
}

func Error(err error) slog.Attr {
	return slog.String(ErrorKey, err.Error())
}

func Project(val string) slog.Attr {
	return slog.String(ProjectKey, val)
}

func Organization(val string) slog.Attr {
	return slog.String(OrganizationKey, val)
}

func HttpAccessLogRecordType() slog.Attr {
	return slog.String(LogRecordTypeKey, HttpAccessLogRecordTypeVal)
}
