package git

import (
	"fmt"
	"os/exec"
)

func RunGit(args ...string) (e error) {
	as := args[1:]
	cmd := exec.Command("git", as...)
	var output []byte
	output, e = cmd.CombinedOutput()
	fmt.Print(string(output))
	return
}
