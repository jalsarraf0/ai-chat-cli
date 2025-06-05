//go:build unit

package shell

import "testing"

func TestFromBaseName(t *testing.T) {
	cases := map[string]string{
		"/bin/bash":          "bash",
		"C:/Windows/cmd.exe": "cmd.exe",
	}
	for in, want := range cases {
		if got := fromBaseName(in); got != want {
			t.Fatalf("%s => %s want %s", in, got, want)
		}
	}
}
