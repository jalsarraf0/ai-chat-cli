// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package config handles configuration loading and saving.
package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	v    = viper.New()
	path string
)

// Reset is intended for tests to reinitialize the package state.
func Reset() {
	v = viper.New()
	path = ""
}

// Load reads configuration from file, env and flags.
func Load(p string) error {
	if p == "" {
		p = defaultPath()
	}
	path = p
	v.SetConfigFile(p)
	v.SetConfigType("yaml")
	v.SetEnvPrefix("AICHAT")
	v.AutomaticEnv()
	if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
		return err
	}
	if err := v.ReadInConfig(); err != nil {
		var e viper.ConfigFileNotFoundError
		if !errors.As(err, &e) && !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}
	return validate()
}

// Save writes configuration to disk.
func Save() error {
	if err := v.WriteConfigAs(path); err != nil {
		return err
	}
	return nil
}

// Set updates a key and saves.
func Set(key string, val any) error {
	v.Set(key, val)
	if err := validate(); err != nil {
		return err
	}
	return Save()
}

// GetString returns a string value.
func GetString(key string) string { return v.GetString(key) }

// GetFloat64 returns a float64 value.
func GetFloat64(key string) float64 { return v.GetFloat64(key) }

// GetInt returns an int value.
func GetInt(key string) int { return v.GetInt(key) }

// defaultPath returns the platform-specific config file path.
func defaultPath() string { return defaultPathImpl() }

var allowedModels = map[string]struct{}{
	"gpt-4":         {},
	"gpt-3.5-turbo": {},
}

func validate() error {
	if k := v.GetString("openai_api_key"); k == "" {
		return errors.New("openai_api_key required")
	}
	if m := v.GetString("model"); m != "" {
		if _, ok := allowedModels[m]; !ok {
			return errors.New("invalid model")
		}
	}
	return nil
}

// Path returns the loaded config path.
func Path() string { return path }
