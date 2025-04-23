package builtins

import (
	"os"
	"os/user"
)

var Builtins = map[string]func([]string) error{
	"cd":    ChangeDir,
	"exit":  Exit,
}

func ChangeDir(args []string) error {
	target := ""
	if len(args) > 1 {
		target = args[1]
	} else {
		usr, err := user.Current()
		if err != nil {
			return err
		}
		target = usr.HomeDir
	}
	return os.Chdir(target)
}

func Exit(args []string) error {
	os.Exit(0)
	return nil
}