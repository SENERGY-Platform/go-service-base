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

package types

import (
	"encoding/json"
	"fmt"
	envldr "github.com/SENERGY-Platform/go-env-loader"
	"reflect"
	"time"
)

type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch val := v.(type) {
	case float64:
		*d = Duration(time.Duration(val))
		return nil
	case string:
		tmp, err := time.ParseDuration(val)
		if err != nil {
			return err
		}
		*d = Duration(tmp)
		return nil
	default:
		return fmt.Errorf("invalid format: %v", val)
	}
}

func DurationEnvTypeParser() (reflect.Type, envldr.Parser) {
	return reflect.TypeOf(Duration(0)), DurationEnvParser
}

func DurationEnvParser(_ reflect.Type, val string, _ []string, _ map[string]string) (interface{}, error) {
	d, err := time.ParseDuration(val)
	if err != nil {
		return Duration(0), err
	}
	return Duration(d), nil
}
