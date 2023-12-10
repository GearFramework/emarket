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

package backend

import (
	"github.com/GearFramework/emarket/internal/app"
	"github.com/GearFramework/emarket/internal/pkg/alog"
	"github.com/GearFramework/emarket/internal/pkg/server"
)

type Backend struct {
	EnvFile string
	Logger  *alog.Alog
	Flags   *Flags
	Config  *app.Config
	Server  *server.HttpServer
}

func NewBackend(envFile string) *Backend {
	return &Backend{
		EnvFile: envFile,
		Logger:  alog.NewLogger(),
	}
}

func (app *Backend) Init() error {
	if err := NewEnv(app.EnvFile); err != nil {
		return err
	}
	app.Flags = GetFlags(GetDefaultFlags())
	app.Server = server.NewServer(NewServerConfig())
	err := app.Server.Init(app.initRoutes)
	return err
}

func (app *Backend) Run() error {
	if err := app.Server.Up(); err != nil {
		return err
	}
	return nil
}

func (app *Backend) Stop() {
}
