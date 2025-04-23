package main

import (
	"whussh/shell"
	"whussh/shell/utils"
)

func main() {
	utils.TrapInterrupt()
	shell.Start()
}