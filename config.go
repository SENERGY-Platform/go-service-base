/*
 * Copyright 2022 InfAI (CC SES)
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

package srv_base

import (
	"encoding/json"
	"github.com/y-du/go-env-loader"
	"github.com/y-du/go-log-level/level"
	"os"
	"reflect"
)

var logLevelParser envldr.Parser = func(t reflect.Type, val string, params []string, kwParams map[string]string) (interface{}, error) {
	return level.Parse(val)
}

var typeParsers = map[reflect.Type]envldr.Parser{
	reflect.TypeOf(level.Off): logLevelParser,
}

func LoadConfig(path *string, cfg any) error {
	if path != nil {
		file, err := os.Open(*path)
		if err != nil {
			return err
		}
		defer file.Close()
		decoder := json.NewDecoder(file)
		if err = decoder.Decode(cfg); err != nil {
			return err
		}
	}
	return envldr.LoadEnvUserParser(cfg, nil, typeParsers, nil)
}
