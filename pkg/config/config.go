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
	v            = viper.New()
	path         string
	skipValidate bool
	// ErrAPIKeyMissing indicates the OpenAI key was not provided.
	ErrAPIKeyMissing = errors.New("openai_api_key required")
)

// Reset is intended for tests to reinitialize the package state.
func Reset() {
	v = viper.New()
	path = ""
	skipValidate = false
}

// SkipValidation controls whether Load skips validation. Intended for CLI config commands.
func SkipValidation(enable bool) { skipValidate = enable }

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
	_ = v.BindEnv("openai_api_key", "OPENAI_API_KEY")
	if err := os.MkdirAll(filepath.Dir(p), 0o700); err != nil {
		return err
	}
	if err := v.ReadInConfig(); err != nil {
		var nf viper.ConfigFileNotFoundError
		var pe viper.ConfigParseError
		switch {
		case errors.As(err, &nf), errors.Is(err, os.ErrNotExist):
			// ignore missing file
		case errors.As(err, &pe):
			// ignore malformed config and continue with defaults
		default:
			return err
		}
	}
	if skipValidate {
		return nil
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

// Get returns a value of any type.
func Get(key string) any { return v.Get(key) }

// IsSet returns true if key exists.
func IsSet(key string) bool { return v.IsSet(key) }

// All returns all configuration as a map.
func All() map[string]any { return v.AllSettings() }

// defaultPath returns the platform-specific config file path.
func defaultPath() string { return defaultPathImpl() }

var allowedModels = map[string]struct{}{
	"gpt-4o":                 {},
	"gpt-4o-mini":            {},
	"gpt-4o-audio-preview":   {},
	"gpt-4o-2024-05-13":      {},
	"gpt-4.1":                {},
	"gpt-4.1-mini":           {},
	"gpt-4.1-nano":           {},
	"gpt-4.1-2025-04-14":     {},
	"gpt-4":                  {},
	"gpt-4-32k":              {},
	"gpt-4-turbo":            {},
	"gpt-4-turbo-preview":    {},
	"gpt-4-vision-preview":   {},
	"gpt-4-0314":             {},
	"gpt-4-0613":             {},
	"gpt-4-0125-preview":     {},
	"gpt-3.5-turbo":          {},
	"gpt-3.5-turbo-16k":      {},
	"gpt-3.5-turbo-0125":     {},
	"gpt-3.5-turbo-1106":     {},
	"text-embedding-3-large": {},
	"text-embedding-3-small": {},
	"text-embedding-ada-002": {},
	"whisper-1":              {},
	"dall-e-3":               {},
	"moderation-latest":      {},
	"moderation-v1":          {},
	"gpt-4o-nano":            {},
	"gpt-image-1":            {},
}

func validate() error {
	if k := v.GetString("openai_api_key"); k == "" {
		return ErrAPIKeyMissing
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
