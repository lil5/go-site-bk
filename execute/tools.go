package execute

import (
	"fmt"
	"os/exec"
)

func runExecf(format string, a ...interface{}) error {
	cmdLine := fmt.Sprintf(format, a...)
	cmd := exec.Command("bash", "-c", cmdLine)
	// cmd.Stdin = strings.NewReader("some input")
	// var out bytes.Buffer
	// cmd.Stdout = &out
	err := cmd.Run()

	return err
}
