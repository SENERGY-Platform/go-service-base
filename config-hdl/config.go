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
	envldr "github.com/SENERGY-Platform/go-env-loader"
	"github.com/SENERGY-Platform/go-service-base/config-hdl/types"
	"io/fs"
	"os"
	"reflect"
	"time"
)

var defaultTypeParsers = map[reflect.Type]envldr.Parser{
	reflect.TypeOf(fs.ModePerm):       fileModeParser,
	reflect.TypeOf(types.Secret("")):  secretStringParser,
	reflect.TypeOf(time.Nanosecond):   durationParser,
	reflect.TypeOf(types.Duration(0)): pkgDurationParser,
}

func Load(cfg any, keywordParsers map[string]envldr.Parser, typeParsers map[reflect.Type]envldr.Parser, kindParsers map[reflect.Kind]envldr.Parser, paths ...string) error {
	for _, p := range paths {
		if p != "" {
			if err := readConfig(p, cfg); err != nil {
				return err
			}
		}
	}
	if len(typeParsers) > 0 {
		for r, parser := range defaultTypeParsers {
			if _, ok := typeParsers[r]; !ok {
				typeParsers[r] = parser
			}
		}
	} else {
		typeParsers = defaultTypeParsers
	}
	return envldr.LoadEnvUserParser(cfg, keywordParsers, typeParsers, kindParsers)
}

func readConfig(path string, cfg any) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	return decoder.Decode(cfg)
}
