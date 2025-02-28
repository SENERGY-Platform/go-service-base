/*
 * Copyright 2025 InfAI (CC SES)
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

package env_parser

import (
	envldr "github.com/SENERGY-Platform/go-env-loader"
	"io/fs"
	"reflect"
	"strconv"
	"time"
)

func FileModeEnvTypeParser() (reflect.Type, envldr.Parser) {
	return reflect.TypeOf(fs.ModePerm), FileModeEnvParser
}

func FileModeEnvParser(_ reflect.Type, val string, _ []string, _ map[string]string) (interface{}, error) {
	fm, err := strconv.ParseInt(val, 8, 32)
	return fs.FileMode(fm), err
}

func DurationEnvTypeParser() (reflect.Type, envldr.Parser) {
	return reflect.TypeOf(time.Nanosecond), DurationEnvParser
}

func DurationEnvParser(_ reflect.Type, val string, _ []string, _ map[string]string) (interface{}, error) {
	return time.ParseDuration(val)
}
