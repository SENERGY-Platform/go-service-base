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

package types

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	envldr "github.com/SENERGY-Platform/go-env-loader"
	"reflect"
)

type Secret string

func (s Secret) Value() string {
	return string(s)
}

func (s Secret) String() string {
	return getRandomStr()
}

func (s Secret) MarshalJSON() ([]byte, error) {
	return json.Marshal(getRandomStr())
}

func (s *Secret) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*s = Secret(str)
	return nil
}

func getRandomStr() string {
	rb := make([]byte, 8)
	_, err := rand.Read(rb)
	if err != nil {
		return err.Error()
	}
	return hex.EncodeToString(rb)
}

func SecretEnvTypeParser() (reflect.Type, envldr.Parser) {
	return reflect.TypeOf(Secret("")), SecretEnvParser
}

func SecretEnvParser(_ reflect.Type, val string, _ []string, _ map[string]string) (interface{}, error) {
	return Secret(val), nil
}
