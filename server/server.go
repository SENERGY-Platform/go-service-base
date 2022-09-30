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

package server

import (
	"context"
	"github.com/SENERGY-Platform/go-service-base/logger"
	"github.com/SENERGY-Platform/go-service-base/types"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	startMsg       = "starting server"
	shutdownMsg    = "server shutdown complete"
	startFailedMsg = "starting server failed: "
)

func handleShutdown(server *http.Server, signals types.SignalSet) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, signals.ToSlice()...)
	go func() {
		sig := <-shutdown
		logger.Logger.Warningf("received signal '%s'", sig)
		logger.Logger.Info("initiating shutdown ...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			logger.Logger.Error("server forced to shutdown: ", err)
		}
	}()
}

func Start(server *http.Server, listener net.Listener, signals types.SignalSet) {
	logger.Logger.Info(startMsg + " ...")
	handleShutdown(server, signals)
	if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
		logger.Logger.Error(startFailedMsg, err)
	} else {
		logger.Logger.Info(shutdownMsg)
	}
}

func StartTLS(server *http.Server, listener net.Listener, signals types.SignalSet, certFile string, keyFile string) {
	logger.Logger.Info(startMsg + " with TLS ...")
	handleShutdown(server, signals)
	if err := server.ServeTLS(listener, certFile, keyFile); err != nil && err != http.ErrServerClosed {
		logger.Logger.Error(startFailedMsg, err)
	} else {
		logger.Logger.Info(shutdownMsg)
	}
}
