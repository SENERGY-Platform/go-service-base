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
	"github.com/SENERGY-Platform/go-service-base/config-hdl/types"
	"os"
	"path"
	"testing"
)

type testConf struct {
	V1 string       `json:"v1"`
	V2 types.Secret `json:"v2"`
}

func TestLoad(t *testing.T) {
	tmpDir := t.TempDir()
	p1 := path.Join(tmpDir, "c1.json")
	f1, err := os.Create(p1)
	if err != nil {
		t.Fatal(err)
	}
	defer f1.Close()
	_, err = f1.WriteString("{\"v1\": \"test1\", \"v2\": \"test2\"}")
	if err != nil {
		t.Fatal(err)
	}
	p2 := path.Join(tmpDir, "c2.json")
	f2, err := os.Create(p2)
	if err != nil {
		t.Fatal(err)
	}
	defer f2.Close()
	_, err = f2.WriteString("{\"v1\": \"test3\"}")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("single conf", func(t *testing.T) {
		a := testConf{
			V1: "test1",
			V2: types.Secret("test2"),
		}
		var b testConf
		err = Load(&b, nil, nil, nil, p1)
		if err != nil {
			t.Error(err)
		}
		if b.V1 != a.V1 {
			t.Errorf("expected %s, got %s", a.V1, b.V1)
		}
		if b.V2.Value() != a.V2.Value() {
			t.Errorf("expected %s, got %s", a.V2.Value(), b.V2.Value())
		}
	})
	t.Run("multiple conf", func(t *testing.T) {
		a := testConf{
			V1: "test3",
			V2: types.Secret("test2"),
		}
		var b testConf
		err = Load(&b, nil, nil, nil, p1, p2)
		if err != nil {
			t.Error(err)
		}
		if b.V1 != a.V1 {
			t.Errorf("expected %s, got %s", a.V1, b.V1)
		}
		if b.V2.Value() != a.V2.Value() {
			t.Errorf("expected %s, got %s", a.V2.Value(), b.V2.Value())
		}
	})
}
