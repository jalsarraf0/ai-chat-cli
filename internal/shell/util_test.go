package shell

import "testing"

func TestClassify(t *testing.T) {
	cases := map[string]Kind{
		"bash":     Bash,
		"zsh":      Zsh,
		"fish":     Fish,
		"pwsh.exe": PowerShell,
		"cmd":      Cmd,
		"unknown":  Unknown,
	}
	for in, want := range cases {
		if got := classify(in); got != want {
			t.Fatalf("%s => %v want %v", in, got, want)
		}
	}
}

func TestKindString(t *testing.T) {
	if Bash.String() != "bash" {
		t.Fatalf("String() mismatch")
	}
}

func TestDetectFromEnv(t *testing.T) {
	got := detectFromEnv("/bin/zsh", "")
	if got != Zsh {
		t.Fatalf("want zsh got %v", got)
	}
	if detectFromEnv("", `C:/Windows/System32/cmd.exe`) != Cmd {
		t.Fatalf("cmd detection failed")
	}
}
