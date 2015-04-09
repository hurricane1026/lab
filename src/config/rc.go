package config

import (
	"fmt"
	"github.com/hurricane1026/gcfg"
	"os"
	"os/user"
)

type LabConfig struct {
	Auth struct {
		Remote   string
		Token    string
		Name     string
		Provider string
	}
}

var Labc LabConfig
var labrc_path string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	if usr, e := user.Current(); e != nil {
		check(e)
	} else {
		labrc_path = fmt.Sprintf("%s%c%s", usr.HomeDir, os.PathSeparator, ".labrc")
	}
	return
}

func Rc_exists() bool {
	if _, e := os.Stat(labrc_path); e == nil {
		return true
	}
	return false
}

func Rc_Load() bool {
	if e := gcfg.ReadFileInto(&Labc, labrc_path); e != nil {
		return false
	}
	return len(Labc.Auth.Name) > 0 && len(Labc.Auth.Remote) > 0 && len(Labc.Auth.Name) > 0 && len(Labc.Auth.Provider) > 0
}

func Rc_Save() {
	template := "[auth]\n\tname = %s\n\tremote = %s\n\ttoken = %s\n\tprovider = %s\n"
	content := fmt.Sprintf(template, Labc.Auth.Name, Labc.Auth.Remote, Labc.Auth.Token, Labc.Auth.Provider)
	f, err := os.Create(labrc_path)
	check(err)
	defer f.Close()
	f.WriteString(content)
	f.Sync()
	return
}

func Rc_Set(remote, token, name, provider string) bool {
	if len(remote) > 0 && len(token) > 0 && len(name) > 0 && len(provider) > 0 {
		Labc.Auth.Remote = remote
		Labc.Auth.Token = token
		Labc.Auth.Name = name
		Labc.Auth.Provider = provider
		return true
	}
	return false
}
