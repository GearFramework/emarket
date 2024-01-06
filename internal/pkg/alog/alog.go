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
)

const LogInfoPattern = "\u001B[0m%s [\u001B[34mINFO\u001B[0m] [\u001B[32m%s\u001B[0m] %s"
const LogWarningPattern = "\u001B[0m%s [\u001B[33mWARNING\u001B[0m] [\u001B[32m%s\u001B[0m] %s"
const LogErrorPattern = "\u001B[0m%s [\u001B[31mERROR\u001B[0m] [\u001B[32m%s\u001B[0m] %s"
const LogFatalPattern = "\u001B[0m%s [\u001B[31mFATAL\u001B[0m] [\u001B[32m%s\u001B[0m] %s"

type Alog struct {
	name string
}

func NewLogger(name string) *Alog {
	return &Alog{name}
}

func (l *Alog) Infof(format string, params ...interface{}) {
	l.Info(fmt.Sprintf(format, params...))
}

func (l *Alog) Info(message string) {
	fmt.Printf(LogInfoPattern+"\n", LogTimeNow(), l.name, message)
}

func (l *Alog) Errorf(format string, params ...interface{}) {
	l.Error(fmt.Sprintf(format, params...))
}

func (l *Alog) Error(message string) {
	fmt.Printf(LogErrorPattern+"\n", LogTimeNow(), l.name, message)
}

func (l *Alog) Warnf(format string, params ...interface{}) {
	l.Warn(fmt.Sprintf(format, params...))
}

func (l *Alog) Warn(message string) {
	fmt.Printf(LogWarningPattern+"\n", LogTimeNow(), l.name, message)
}

func (l *Alog) Fatalf(format string, params ...interface{}) {
	l.Fatal(fmt.Sprintf(format, params...))
}

func (l *Alog) Fatal(message string) {
	fmt.Printf(LogFatalPattern+"\n", LogTimeNow(), l.name, message)
}
