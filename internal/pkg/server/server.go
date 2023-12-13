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

package server

import (
	"github.com/GearFramework/emarket/internal/pkg/alog"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpServer struct {
	HTTP   *http.Server
	Router *gin.Engine
	Logger *alog.Alog
	Config *Config
}

type MiddlewareFunc func() gin.HandlerFunc

func NewServer(conf *Config) *HttpServer {
	gin.SetMode(gin.ReleaseMode)
	return &HttpServer{
		Config: conf,
		Logger: alog.NewLogger(),
		Router: gin.New(),
	}
}

func (serv *HttpServer) SetMiddleware(mw MiddlewareFunc) *HttpServer {
	serv.Router.Use(mw())
	return serv
}

func (serv *HttpServer) Init(initRoutes func()) error {
	initRoutes()
	return nil
}

func (serv *HttpServer) Up() error {
	serv.HTTP = &http.Server{
		Addr:    serv.Config.GetDSN(),
		Handler: serv.Router,
	}
	serv.Logger.Infof("Start server at the %s", serv.Config.GetDSN())
	err := serv.HTTP.ListenAndServe()
	if err != nil {
		serv.Logger.Errorf("Failed to Listen and Serve %s", err.Error())
		return err
	}
	return nil
}
