/*
 * Copyright 2020 zpxio (Jeff Sharpe)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetHandler(text.New(os.Stderr))

	log.Infof("Starting up...")

	viper.SetConfigName("config")

	log.Infof("Initializing configuration")

	// Initialize flags
	pflag.Int("dashboard-port", 80, "The port for unencrypted dashboard connections")
	pflag.IP("dashboard-host", net.IPv4(0, 0, 0, 0), "The host IP to listen on for dashboard connections")

	// Parse flags
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	// Load configuration from files
	configErr := viper.ReadInConfig()
	if configErr != nil {
		if _, ok := configErr.(viper.ConfigFileNotFoundError); ok {
			log.Infof("No configuration files found.\n")
		} else {
			log.Fatalf("Could not read config file: %s\n", configErr)
		}
	}

	// Apply flags
	flagErr := viper.BindPFlags(pflag.CommandLine)
	if flagErr != nil {
		log.Infof("Error while translating flags: %s", flagErr)
	}

	// Register actions

	// Initialize the event system

	// Start the monitors

	// Start

	// Set up signal monitoring
	termSignals := make(chan os.Signal, 1)
	exitChan := make(chan bool, 1)
	signal.Notify(termSignals, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		shutdownSignal := <-termSignals

		log.Infof("Received shutdown signal: %s", shutdownSignal)
		exitChan <- true
	}()

	// Wait for shutdown signals
	<-exitChan
	log.Info("Initiating shutdown.")
}
