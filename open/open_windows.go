//go:build windows
// +build windows

package open

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

var (
	cmd      = "url.dll,FileProtocolHandler"
	runDll32 = filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")
)

func open(input string) error {
	validInput := regexp.MustCompile(`^[a-zA-Z0-9_\-\./\\:?=&#%]+$`)
	if !validInput.MatchString(input) {
		return fmt.Errorf("invalid input")
	}
	return exec.Command(runDll32, cmd, input).Run()
}
