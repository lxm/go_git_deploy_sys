package git

import (
	"fmt"
	//"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func RunCmd(cmdName string, cmdArgs []string) {
	cmd := exec.Command(cmdName, cmdArgs...)
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output) // => go version go1.3 darwin/amd64
}
func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}
func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}
func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}
