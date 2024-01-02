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
	"fmt"
	"github.com/GearFramework/emarket/internal/pkg/gear"
	"os"
)

const (
	DefaultAddr = "localhost"
	DefaultPort = 8080
)

type Config struct {
	Addr string
	Port int
}

func NewServerConfig() *Config {
	return &Config{
		Addr: gear.Getenv("BACKEND_ADDR", DefaultAddr),
		Port: gear.AtoI(os.Getenv("BACKEND_PORT"), DefaultPort),
	}
}

func (conf *Config) GetDSN() string {
	return fmt.Sprintf("%s:%d", conf.Addr, conf.Port)
}
