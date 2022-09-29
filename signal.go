/*
 * Copyright 2022 InfAI (CC SES)
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

package srv_base

import (
	"os"
	"syscall"
)

type ShutdownSignals map[os.Signal]struct{}

func (s ShutdownSignals) Add(sig os.Signal) {
	s[sig] = struct{}{}
}

func (s ShutdownSignals) Remove(sig os.Signal) {
	delete(s, sig)
}

func (s ShutdownSignals) Has(sig os.Signal) bool {
	if _, ok := s[sig]; ok {
		return true
	}
	return false
}

func (s ShutdownSignals) ToSlice() (sl []os.Signal) {
	for sig := range s {
		sl = append(sl, sig)
	}
	return sl
}

var DefaultSignals = ShutdownSignals{
	syscall.SIGINT:  struct{}{},
	syscall.SIGTERM: struct{}{},
}
