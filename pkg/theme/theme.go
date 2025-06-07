package theme

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/jalsarraf0/ai-chat-cli/pkg/embedutil"
)

// Palette represents a colour palette.
type Palette struct {
	Background string `json:"background"`
}

// Load reads the palette with optional name. When name is empty it selects a
// default based on COLORTERM ("light" selects the light palette).
func Load(name string) Palette {
	if name == "" {
		ct := strings.ToLower(os.Getenv("COLORTERM"))
		if strings.Contains(ct, "light") {
			name = "themes/light.json"
		} else {
			name = "themes/dark.json"
		}
	}
	data, err := embedutil.Read(name)
	if err != nil {
		data = []byte(`{"background":""}`)
	}
	var p Palette
	_ = json.Unmarshal(data, &p)
	return p
}
