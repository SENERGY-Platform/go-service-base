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
	"encoding/json"
	"testing"
)

func TestSecret_MarshalJSON(t *testing.T) {
	secret := Secret("test")
	b, err := json.Marshal(secret)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) == "test" {
		t.Errorf("%s should not be equal to test", string(b))
	}
}

func TestSecret_UnmarshalJSON(t *testing.T) {
	var test struct {
		S Secret `json:"s"`
	}
	err := json.Unmarshal([]byte("{\"s\": \"test\"}"), &test)
	if err != nil {
		t.Fatal(err)
	}
	if string(test.S) != "test" {
		t.Errorf("%s should be equal to test", string(test.S))
	}
}

func TestSecret_String(t *testing.T) {
	secret := Secret("test")
	if secret.String() == "test" {
		t.Errorf("%s should not be equal to test", string(secret))
	}
}
