//go:build tools
// +build tools

package tools

import (
    _ "github.com/spf13/cobra/doc"
    _ "github.com/cpuguy83/go-md2man/v2"
    _ "github.com/russross/blackfriday/v2"
    _ "gopkg.in/yaml.v3"
)

