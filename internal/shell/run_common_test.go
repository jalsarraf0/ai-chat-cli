package shell

import (
	"context"
	"runtime"
	"testing"
)

func TestRunCmd_LookPathStub(t *testing.T) {
	t.Parallel()

	lookPathMu.Lock()
	orig := lookPath
	lookPath = func(file string) (string, error) {
		if file == "echo" {
			if runtime.GOOS == "windows" {
				// cmd's internal 'echo' isn't an external binary,
				// so use cmd.exe and adapt the argument list below.
				return "C:\\Windows\\System32\\cmd.exe", nil
			}
			// Unix
			return "/bin/echo", nil
		}
		// Delegate to the original function for other binaries to
		// avoid interfering with parallel tests.
		return orig(file)
	}
	lookPathMu.Unlock()
	defer func() {
		lookPathMu.Lock()
		lookPath = orig
		lookPathMu.Unlock()
	}()

	var (
		stdout string
		stderr string
		err    error
	)

	if runtime.GOOS == "windows" {
		stdout, stderr, err = runCmd(context.Background(),
			"echo",
			[]string{"/C", "echo", "hi"},
		)
	} else {
		stdout, stderr, err = runCmd(context.Background(), "echo", []string{"hi"})
	}

	if err != nil {
		t.Fatalf("runCmd failed: %v", err)
	}

	want := "hi\n"
	if runtime.GOOS == "windows" {
		// cmd.exe adds trailing CRLF by default
		want = "hi\r\n"
	}
	if stdout != want || stderr != "" {
		t.Fatalf("unexpected output: %q / %q", stdout, stderr)
	}
}
