package config

import (
	"os"
	"testing"
)

func Test_Init(t *testing.T) {
	return
}

func Test_Exist(t *testing.T) {
	if Rc_exists() == false {
		t.Error(".labrc should exist")
	}
}

func Test_Validation(t *testing.T) {
	if Rc_Load() == false {
		t.Log(Labc)
		t.Error("Labc is not conrecct")
	}
}

func Test_Smoke(t *testing.T) {
	os.Remove(labrc_path)
	Rc_Set("test.test.com", "asdg123412", "testname", "gitlab")
	Rc_Save()
	Labc.Auth.Name = ""
	Labc.Auth.Token = ""
	Labc.Auth.Remote = ""
	Rc_Load()
	if Labc.Auth.Name != "testname" || Labc.Auth.Token != "asdg123412" ||
		Labc.Auth.Remote != "test.test.com" {
		t.Error("Smoke failed")
	}
}
