// Package assets contains embedded resources.
package assets

import "embed"

// FS embeds default templates and colour themes.
//
//go:embed templates/*.tmpl themes/*.json
var FS embed.FS
