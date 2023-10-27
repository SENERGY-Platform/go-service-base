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
	"errors"
	"github.com/SENERGY-Platform/go-service-base/job-hdl/lib"
	"sync"
	"time"
)

type job struct {
	mu    sync.RWMutex
	tFunc func(context.Context, context.CancelFunc) error
	ctx   context.Context
	cFunc context.CancelFunc
	lib.Job
}

func (j *job) CallTarget(cbk func()) {
	Logger.Debugf("job '%s' starting ...", j.ID)
	j.mu.Lock()
	t := time.Now().UTC()
	j.Started = &t
	j.mu.Unlock()
	err := j.tFunc(j.ctx, j.cFunc)
	j.mu.Lock()
	if err != nil {
		j.Error = &lib.JobErr{
			Message: err.Error(),
			Code:    ErrCodeMapper(err),
		}
		Logger.Warningf("job '%s' got error: %s", j.ID, err.Error())
	}
	t2 := time.Now().UTC()
	j.Completed = &t2
	j.mu.Unlock()
	Logger.Debugf("job '%s' completed", j.ID)
	cbk()
}

func (j *job) IsCanceled() bool {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return errors.Is(j.ctx.Err(), context.Canceled)
}

func (j *job) Cancel() {
	j.cFunc()
	j.mu.Lock()
	t := time.Now().UTC()
	j.Canceled = &t
	j.mu.Unlock()
}

func (j *job) Meta() lib.Job {
	j.mu.RLock()
	defer j.mu.RUnlock()
	return j.Job
}
