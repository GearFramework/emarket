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

package auth

import (
	"github.com/GearFramework/emarket/internal/app"
	"github.com/GearFramework/emarket/internal/models"
	"github.com/GearFramework/emarket/internal/pkg/alog"
	"github.com/GearFramework/emarket/internal/pkg/cache/redis"
	"github.com/GearFramework/emarket/internal/pkg/server"
	"github.com/GearFramework/emarket/internal/pkg/server/middleware"
	"github.com/gin-gonic/gin"
)

type ServiceAuth struct {
	EnvFile string
	Flags   *Flags
	Server  *server.HttpServer
	Config  *app.ServiceAuthConfig
	Cache   models.Cachable
	Storage models.Storable
	logger  *alog.Alog
}

func NewServiceAuth(envFile string) *ServiceAuth {
	return &ServiceAuth{
		EnvFile: envFile,
		logger:  alog.NewLogger(),
	}
}

func (app *ServiceAuth) Run() error {
	if err := app.Server.Up(); err != nil {
		return err
	}
	return nil
}

// Init initialize auth service
func (app *ServiceAuth) Init() error {
	if err := NewEnv(app.EnvFile); err != nil {
		return err
	}
	app.Flags = GetFlags(GetDefaultFlags())
	app.Config = NewAuthConfig()
	app.Cache = redis.NewCache(redis.NewRedisConfig())
	if err := app.Cache.(*redis.CacheRedis).InitCache(); err != nil {
		return err
	}
	//app.Storage = db.NewStorage()
	app.Server = server.NewServer(server.NewServerConfig())
	app.Server.SetMiddleware(func() gin.HandlerFunc {
		return middleware.Logger()
	})
	err := app.Server.Init(app.initRoutes)
	return err
}

func (app *ServiceAuth) Stop() {
}

func (app *ServiceAuth) Logger() *alog.Alog {
	return app.logger
}
