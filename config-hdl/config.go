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
	"os"
	"reflect"
)

type EnvParser = envldr.Parser

type EnvKeywordParser func() (string, EnvParser)

type EnvTypeParser func() (reflect.Type, EnvParser)

type EnvKindParser func() (reflect.Kind, EnvParser)

func Load(cfg any, envKeywordParsers []EnvKeywordParser, envTypeParsers []EnvTypeParser, envKindParsers []EnvKindParser, paths ...string) error {
	for _, p := range paths {
		if p != "" {
			if err := readConfig(p, cfg); err != nil {
				return err
			}
		}
	}
	keywordParserMap := make(map[string]EnvParser)
	for _, keywordParser := range envKeywordParsers {
		kw, p := keywordParser()
		keywordParserMap[kw] = p
	}
	typeParserMap := make(map[reflect.Type]EnvParser)
	for _, typeParser := range envTypeParsers {
		t, p := typeParser()
		typeParserMap[t] = p
	}
	kindParserMap := make(map[reflect.Kind]EnvParser)
	for _, kindParser := range envKindParsers {
		k, p := kindParser()
		kindParserMap[k] = p
	}
	return envldr.LoadEnvUserParser(cfg, keywordParserMap, typeParserMap, kindParserMap)
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
