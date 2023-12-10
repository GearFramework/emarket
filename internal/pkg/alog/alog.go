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

package alog

import (
	"fmt"
	"os"
)

type Alog struct{}

func NewLogger() *Alog {
	return &Alog{}
}

func (l *Alog) Infof(format string, params ...interface{}) {
	l.Info(fmt.Sprintf(format, params...))
}

func (l *Alog) Info(message string) {
	fmt.Printf("%s INFO %s\n", LogTimeNow(), message)
}

func (l *Alog) Errorf(format string, params ...interface{}) {
	l.Info(fmt.Sprintf(format, params...))
}

func (l *Alog) Error(message string) {
	fmt.Printf("%s%s ERROR %s%s\n", string(ColorError), LogTimeNow(), message, string(ColorReset))
}

func (l *Alog) Warnf(format string, params ...interface{}) {
	l.Info(fmt.Sprintf(format, params...))
}

func (l *Alog) Warn(message string) {
	fmt.Printf("%s%s ERROR %s%s\n", string(ColorWarning), LogTimeNow(), message, string(ColorReset))
}

func (l *Alog) Fatalf(format string, params ...interface{}) {
	l.Info(fmt.Sprintf(format, params...))
}

func (l *Alog) Fatal(message string) {
	fmt.Printf("%s%s FATAL %s%s\n", string(ColorError), LogTimeNow(), message, string(ColorReset))
	os.Exit(1)
}
