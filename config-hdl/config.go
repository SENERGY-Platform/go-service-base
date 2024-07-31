/*
 * Copyright 2024 InfAI (CC SES)
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

package config_hdl

import (
	"encoding/json"
	"github.com/SENERGY-Platform/go-service-base/config-hdl/types"
	envldr "github.com/y-du/go-env-loader"
	"io/fs"
	"os"
	"reflect"
	"strconv"
)

var fileModeParser = func(t reflect.Type, val string, params []string, kwParams map[string]string) (interface{}, error) {
	fm, err := strconv.ParseInt(val, 8, 32)
	return fs.FileMode(fm), err
}

var secretStringParser = func(t reflect.Type, val string, params []string, kwParams map[string]string) (interface{}, error) {
	return types.ParseSecret(val)
}

var defaultTypeParsers = map[reflect.Type]envldr.Parser{
	reflect.TypeOf(fs.ModePerm):      fileModeParser,
	reflect.TypeOf(types.Secret("")): secretStringParser,
}

func Load(path string, cfg any, envKeywordParsers map[string]envldr.Parser, envTypeParsers map[reflect.Type]envldr.Parser, envKindParsers map[reflect.Kind]envldr.Parser) error {
	if path != "" {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(cfg); err != nil {
			return err
		}
	}
	if len(envTypeParsers) > 0 {
		for r, parser := range defaultTypeParsers {
			if _, ok := envTypeParsers[r]; !ok {
				envTypeParsers[r] = parser
			}
		}
	} else {
		envTypeParsers = defaultTypeParsers
	}
	return envldr.LoadEnvUserParser(cfg, envKeywordParsers, envTypeParsers, envKindParsers)
}
