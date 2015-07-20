package command

import (
	"config"
	"fmt"
	"providers"
)

func forkUsage() {
	fmt.Println("lab fork usage: lab fork repo local_directory_name")
}

func forkCommand(args ...string) {
	if len(args) < 2 || args[1] == "help" || args[1] == "h" {
		forkUsage()
	} else {
		if has_config := config.Rc_Load(); !has_config {
			fmt.Println("lab fork usage: please setup your gitlab remote address at first")
			return
		}
		father_repo := args[1]
		scm := providers.NewScm("lab", config.Labc.Auth.Remote, config.Labc.Auth.Token)
		scm.ForkProject(father_repo)

	}
	return
}
