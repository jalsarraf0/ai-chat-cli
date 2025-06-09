# Terminal UI

The `ai-chat tui` command launches an interactive chat interface.

| Key | Action |
|-----|-------|
| `PgUp` | Scroll up |
| `PgDn` | Scroll down |
| `Ctrl+C`/`Esc` | Quit |
| `:q` + `Enter` | Quit |

The UI adapts to terminal size and automatically picks a light or dark theme
based on `$COLORTERM`. Set `NO_COLOR=1` to disable ANSI styling.

![tui session](https://asciinema.org/a/placeholder.svg)
