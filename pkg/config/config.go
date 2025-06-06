// Copyright (c) 2025 AI Chat
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
