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
	ErrorKey         = "error"
	ProjectKey       = "project"
	OrganizationKey  = "organization"
	LogRecordTypeKey = "log_record_type"
)

const (
	HttpAccessLogRecordTypeVal = "http_access"
)

func Error(err error) slog.Attr {
	return slog.String(ErrorKey, err.Error())
}
