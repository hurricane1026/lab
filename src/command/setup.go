package command

import (
	"config"
	"fmt"
	"providers"
)

func setupUsage() {
	fmt.Println("lab setup usage: lab setup gitlab/github url token")
}

func setupCommand(args ...string) {
	//lab setup gitlab https://git.xiaojukeji.com/api/v3 KXGLy5yctc7sUkr4oshh
	if len(args) < 4 || args[1] == "help" || args[1] == "h" {
		setupUsage()
	} else {
		provider := args[1]
		url := args[2]
		token := args[3]
		if provider == "gitlab" {
			provider = "lab"
		}
		scm := providers.NewScm("lab", url, token)
		if user, e := scm.CurrentUser(); e == nil {
			config.Rc_Set(url, token, user, provider)
			config.Rc_Save()
		}
	}
}
