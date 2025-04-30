/*
 * Copyright 2024 InfAI (CC SES)
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

package srv_info_hdl

import (
	"fmt"
	"runtime"
	"time"
)

type Handler struct {
	name    string
	version string
	started time.Time
}

func New(name, version string) *Handler {
	return &Handler{
		name:    name,
		version: version,
		started: time.Now(),
	}
}

func (h *Handler) ServiceInfo() ServiceInfo {
	upTime := time.Since(h.started)
	var mStats runtime.MemStats
	runtime.ReadMemStats(&mStats)
	return ServiceInfo{
		Name:      h.name,
		Version:   h.version,
		UpTime:    upTime.String(),
		UpTimeNs:  upTime.Nanoseconds(),
		MemAlloc:  byteToString(mStats.Alloc),
		MemAllocB: mStats.Alloc,
	}
}

func (h *Handler) Name() string {
	return h.name
}

func (h *Handler) Version() string {
	return h.version
}

func (h *Handler) UpTime() time.Duration {
	return time.Since(h.started)
}

func (h *Handler) MemAlloc() uint64 {
	var mStats runtime.MemStats
	runtime.ReadMemStats(&mStats)
	return mStats.Alloc
}

const unit = 1024

// byteToString
// Source: https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
// Title: Format byte size as kilobytes, megabytes, gigabytes, ...
// Author: Stefan Nilsson
// Licence: https://creativecommons.org/licenses/by/3.0/
// Adaptation: Renamed function, changed string format and defined const variable outside function block.
func byteToString(b uint64) string {
	if b < unit {
		return fmt.Sprintf("%dB", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
