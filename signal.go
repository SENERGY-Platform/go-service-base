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

type SignalSet map[os.Signal]struct{}

func NewSignalSet(signals ...os.Signal) (stSigs SignalSet) {
	stSigs = make(SignalSet)
	if signals != nil && len(signals) > 0 {
		for _, sig := range signals {
			stSigs.Add(sig)
		}
	}
	return
}

func (s SignalSet) Add(sig os.Signal) {
	s[sig] = struct{}{}
}

func (s SignalSet) Remove(sig os.Signal) {
	delete(s, sig)
}

func (s SignalSet) Has(sig os.Signal) bool {
	if _, ok := s[sig]; ok {
		return true
	}
	return false
}

func (s SignalSet) ToSlice() (sl []os.Signal) {
	for sig := range s {
		sl = append(sl, sig)
	}
	return sl
}

var DefaultShutdownSignals = NewSignalSet(syscall.SIGINT, syscall.SIGTERM)
