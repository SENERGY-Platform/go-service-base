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
	"context"
	"github.com/SENERGY-Platform/go-service-base/srv-base/types"
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

func handleShutdown(server *http.Server, signals srv_base_types.SignalSet) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, signals.ToSlice()...)
	go func() {
		sig := <-shutdown
		Logger.Warningf("received signal '%s'", sig)
		Logger.Info("initiating shutdown ...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			Logger.Error("server forced to shutdown: ", err)
		}
	}()
}

func StartServer(server *http.Server, listener net.Listener, signals srv_base_types.SignalSet) {
	Logger.Info(startMsg + " ...")
	handleShutdown(server, signals)
	if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
		Logger.Error(startFailedMsg, err)
	} else {
		Logger.Info(shutdownMsg)
	}
}

func StartServerTLS(server *http.Server, listener net.Listener, signals srv_base_types.SignalSet, certFile string, keyFile string) {
	Logger.Info(startMsg + " with TLS ...")
	handleShutdown(server, signals)
	if err := server.ServeTLS(listener, certFile, keyFile); err != nil && err != http.ErrServerClosed {
		Logger.Error(startFailedMsg, err)
	} else {
		Logger.Info(shutdownMsg)
	}
}
