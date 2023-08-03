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

package srv_base

import (
	"context"
	log_level "github.com/y-du/go-log-level"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Watchdog struct {
	signals map[os.Signal]struct{}
	sigChan chan os.Signal
	hltChan chan struct{}
	ticker  *time.Ticker
	hltFunc []func() bool
	stpFunc []func() error
}

func NewWatchdog(signals ...os.Signal) *Watchdog {
	sig := make(map[os.Signal]struct{})
	for _, s := range signals {
		sig[s] = struct{}{}
	}
	return &Watchdog{
		signals: sig,
		sigChan: make(chan os.Signal, 1),
		hltChan: make(chan struct{}),
	}
}

func (w *Watchdog) RegisterHealthFunc(f func() bool) {
	w.hltFunc = append(w.hltFunc, f)
}

func (w *Watchdog) RegisterStopFunc(f func() error) {
	w.stpFunc = append(w.stpFunc, f)
}

func (w *Watchdog) Start(logger *log_level.Logger) {
	for sig := range w.signals {
		signal.Notify(w.sigChan, sig)
	}
	ctx, cf := context.WithCancel(context.Background())
	w.ticker = time.NewTicker(time.Second)
	defer w.ticker.Stop()
	for _, f := range w.hltFunc {
		w.startHealthcheck(ctx, f)
	}
	select {
	case sig := <-w.sigChan:
		logger.Warningf("received signal '%s'", sig)
		break
	case <-w.hltChan:
		break
	}
	cf()
	w.shutdown(logger)
}

func (w *Watchdog) startHealthcheck(ctx context.Context, f func() bool) {
	go func() {
		for {
			select {
			case <-w.ticker.C:
				if !f() {
					close(w.hltChan)
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (w *Watchdog) shutdown(logger *log_level.Logger) {
	logger.Warning("initiating shutdown ...")
	var wg sync.WaitGroup
	wg.Add(len(w.stpFunc))
	for _, f := range w.stpFunc {
		fu := f
		go func() {
			err := fu()
			if err != nil {
				logger.Error(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
