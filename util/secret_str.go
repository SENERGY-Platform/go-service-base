/*
 * Copyright 2023 InfAI (CC SES)
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

package util

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
)

type SecretString string

func (s SecretString) String() string {
	return string(s)
}

func (s SecretString) MarshalJSON() ([]byte, error) {
	rb := make([]byte, 8)
	_, err := rand.Read(rb)
	if err != nil {
		return nil, err
	}
	return json.Marshal(hex.EncodeToString(rb))
}

func (s *SecretString) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*s = SecretString(str)
	return nil
}
