/*
 * Copyright 2022 zpxio (Jeff Sharpe)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package action

import (
	"fmt"
	"strconv"
)

type ParamType int

const (
	ParamInt    ParamType = 0
	ParamFloat            = 1
	ParamString           = 2
	ParamBool             = 3
)

var paramTypeNames = map[ParamType]string{
	ParamInt:    "int",
	ParamFloat:  "float",
	ParamString: "string",
	ParamBool:   "bool",
}

const (
	DefaultInt    = 0
	DefaultFloat  = 0.0
	DefaultString = ""
	DefaultBool   = false
)

type Param struct {
	Name string
	Type ParamType
	Data interface{}
}

type Action struct {
	Type   string           `json:"type"`
	Name   string           `json:"name"`
	Params map[string]Param `json:"params"`
}

func CreateAction(name string, actionType string) Action {

	data := make(map[string]Param)

	return Action{
		Type:   actionType,
		Name:   name,
		Params: data,
	}
}

func (a Action) GetInt(name string) (int64, error) {
	p, ok := a.Params[name]
	if !ok {
		return DefaultInt, fmt.Errorf("could not find param %s", name)
	}

	switch p.Type {
	case ParamInt:
		return p.Data.(int64), nil
	case ParamFloat:
		return int64(p.Data.(float64)), nil
	case ParamBool:
		if p.Data.(bool) {
			return 1, nil
		}
		return 0, nil
	case ParamString:
		fv, parseError := strconv.ParseInt(p.Data.(string), 10, 64)

		if parseError != nil {
			return DefaultInt, fmt.Errorf("failed to parse error as floating point: %s", parseError.Error())
		}
		return fv, nil
	default:
		return DefaultInt, fmt.Errorf("unsupported param type %s", paramTypeNames[p.Type])
	}
}

func (a Action) GetFloat(name string) (float64, error) {
	p, ok := a.Params[name]
	if !ok {
		return DefaultFloat, fmt.Errorf("could not find param %s", name)
	}

	switch p.Type {
	case ParamInt:
		return float64(p.Data.(int)), nil
	case ParamFloat:
		return p.Data.(float64), nil
	case ParamBool:
		if p.Data.(bool) {
			return 1.0, nil
		}
		return 0.0, nil
	case ParamString:
		fv, parseError := strconv.ParseFloat(p.Data.(string), 64)

		if parseError != nil {
			return DefaultFloat, fmt.Errorf("failed to parse error as floating point: %s", parseError.Error())
		}
		return fv, nil
	default:
		return DefaultInt, fmt.Errorf("unsupported param type %s", paramTypeNames[p.Type])
	}
}

func (a Action) GetBoolean(name string) (bool, error) {
	p, ok := a.Params[name]
	if !ok {
		return DefaultBool, fmt.Errorf("could not find param %s", name)
	}

	switch p.Type {
	case ParamInt:
		if p.Data.(int) == 0 {
			return false, nil
		}
		return true, nil
	case ParamFloat:
		if p.Data.(float64) == 0.000 {
			return false, nil
		}
		return true, nil
	case ParamBool:
		return p.Data.(bool), nil
	case ParamString:
		switch p.Data.(string) {
		case "true", "on", "enabled":
			return true, nil
		case "false", "off", "disabled":
			return false, nil
		default:
			return false, nil
		}
	default:
		return DefaultBool, fmt.Errorf("unsupported param type %s", paramTypeNames[p.Type])
	}
}

func (a Action) GetString(name string) (string, error) {
	p, ok := a.Params[name]
	if !ok {
		return DefaultString, fmt.Errorf("could not find param %s", name)
	}

	switch p.Type {
	case ParamInt:
		return fmt.Sprintf("%d", p.Data.(int64)), nil
	case ParamFloat:
		return fmt.Sprintf("%f", p.Data.(float64)), nil
	case ParamBool:
		if p.Data.(bool) {
			return "true", nil
		}
		return "false", nil
	case ParamString:
		return p.Data.(string), nil
	default:
		return DefaultString, fmt.Errorf("unsupported param type %s", paramTypeNames[p.Type])
	}
}

func (a Action) GetIntWithDefault(name string, defaultValue int64) (int64, error) {
	_, ok := a.Params[name]
	if !ok {
		return defaultValue, nil
	}

	return a.GetInt(name)
}

func (a Action) GetFloatWithDefault(name string, defaultValue float64) (float64, error) {
	_, ok := a.Params[name]
	if !ok {
		return defaultValue, nil
	}

	return a.GetFloat(name)
}
func (a Action) GetBooleanWithDefault(name string, defaultValue bool) (bool, error) {
	_, ok := a.Params[name]
	if !ok {
		return defaultValue, nil
	}

	return a.GetBoolean(name)
}
func (a Action) GetStringWithDefault(name string, defaultValue string) (string, error) {
	_, ok := a.Params[name]
	if !ok {
		return defaultValue, nil
	}

	return a.GetString(name)
}
