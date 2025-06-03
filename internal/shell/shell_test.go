package shell

import (
	"context"
	"testing"
)

func TestDetect(t *testing.T) {
	cases := []struct {
		name     string
		env      map[string]string
		wantName string
		wantErr  bool
	}{
		{"bash", map[string]string{"SHELL": "/bin/bash"}, "bash", false},
		{"zsh", map[string]string{"SHELL": "/usr/bin/zsh"}, "zsh", false},
		{"fish", map[string]string{"SHELL": "/usr/bin/fish"}, "fish", false},
		{"powershell", map[string]string{"PSExecutionPolicy": "RemoteSigned", "SHELL": "pwsh"}, "powershell", false},
		{"cmd", map[string]string{"ComSpec": `C:\\Windows\\System32\\cmd.exe`}, "cmd", false},
		{"unknown", map[string]string{}, "", true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.env {
				t.Setenv(k, v)
			}
			name, path, err := Detect()
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if name != tc.wantName {
				t.Errorf("name=%s want %s", name, tc.wantName)
			}
			if path == "" {
				t.Errorf("path empty")
			}
		})
	}
}

func TestRun(t *testing.T) {
	t.Setenv("SHELL", "/bin/bash")
	stdout, stderr, err := Run(context.Background(), "echo hi")
	if err != nil {
		t.Fatalf("run err: %v stderr: %s", err, stderr)
	}
	if string(stdout) != "hi\n" {
		t.Errorf("stdout=%q", stdout)
	}
	if len(stderr) != 0 {
		t.Errorf("stderr=%q", stderr)
	}
}

func TestRunUnsupported(t *testing.T) {
	t.Setenv("SHELL", "")
	t.Setenv("ComSpec", "")
	t.Setenv("PSExecutionPolicy", "")
	_, _, err := Run(context.Background(), "echo hi")
	if err == nil {
		t.Fatal("expected error")
	}
}
