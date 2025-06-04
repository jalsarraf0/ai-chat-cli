package shell

import (
	"runtime"
	"testing"
)

func TestDetectUnix(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("unix only")
	}
	tests := []struct {
		env  string
		want Kind
	}{
		{"/bin/bash", Bash},
		{"/usr/bin/zsh", Zsh},
		{"/usr/bin/fish", Fish},
	}
	for _, tt := range tests {
		t.Setenv("SHELL", tt.env)
		if got := Detect(); got != tt.want {
			t.Fatalf("env %s want %v got %v", tt.env, tt.want, got)
		}
	}
}
