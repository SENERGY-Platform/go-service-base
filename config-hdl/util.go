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

package config_hdl

import "reflect"

const jsonTagKey = "json"

func StructToMap(o any, jsonKey bool) map[string]any {
	oValue := reflect.ValueOf(o)
	if oValue.Kind() == reflect.Ptr {
		oValue = oValue.Elem()
	}
	oMap := make(map[string]any)
	for i := 0; i < oValue.NumField(); i++ {
		var val any
		if oValue.Field(i).Kind() == reflect.Struct {
			val = StructToMap(oValue.Field(i).Interface(), jsonKey)
		} else {
			val = oValue.Field(i).Interface()
		}
		key := oValue.Type().Field(i).Name
		if jsonKey {
			if tmp, ok := oValue.Type().Field(i).Tag.Lookup(jsonTagKey); ok {
				key = tmp
			}
		}
		oMap[key] = val
	}
	return oMap
}
