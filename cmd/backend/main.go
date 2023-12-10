/*
	Copyright 2023 GearTeam

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package main

import (
	"github.com/GearFramework/emarket/internal/backend"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	app := backend.NewBackend("./.env")
	if err := app.Init(); err != nil {
		return err
	}
	gracefulStop(app.Stop)
	if err := app.Run(); err != nil {
		return err
	}
	return nil
}

func gracefulStop(stopCallback func()) {
	gracefulStopChan := make(chan os.Signal, 1)
	signal.Notify(
		gracefulStopChan,
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	go func() {
		sig := <-gracefulStopChan
		stopCallback()
		log.Printf("Caught sig: %+v\n", sig)
		log.Println("Application graceful stop!")
		os.Exit(0)
	}()
}
