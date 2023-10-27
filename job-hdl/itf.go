/*
 * Copyright 2023 InfAI (CC SES)
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

package job_hdl

import (
	"context"
	"github.com/SENERGY-Platform/go-service-base/job-hdl/lib"
)

type JobHandler interface {
	Create(ctx context.Context, desc string, tFunc func(context.Context, context.CancelFunc) error) (string, error)
	Get(ctx context.Context, id string) (lib.Job, error)
	Cancel(ctx context.Context, id string) error
	List(ctx context.Context, filter lib.JobFilter) ([]lib.Job, error)
	PurgeJobs(ctx context.Context, maxAge int64) (int, error)
}

var Logger interface {
	Warningf(format string, arg ...any)
	Debugf(format string, arg ...any)
}

var (
	ErrCodeMapper  func(error) int
	NewInternalErr func(error) error
	NewNotFoundErr func(error) error
)
