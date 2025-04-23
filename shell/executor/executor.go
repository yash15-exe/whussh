package executor

import (
	"fmt"
	"os"
	"os/exec"
	"whussh/shell/executor/builtins" 
)

func Execute(commands [][]string) {
	if len(commands) == 0 {
		return 
	}

	if handleBuiltin(commands[0]) {
		return
	}

	if len(commands) > 1 {
		executePipeline(commands)
	} else {
		executeSingleCommand(commands[0])
	}
}

func executeSingleCommand(args []string) {

	outputFile := ""
	for i, arg := range args {
		if arg == ">" && i < len(args)-1 {
			outputFile = args[i+1]
			args = args[:i] 
			break
		}
	}

	cmd := exec.Command(args[0], args[1:]...) 


	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("whussh: %v\n", err)
			return
		}
		defer file.Close()
		cmd.Stdout = file
	} else {
		cmd.Stdout = os.Stdout
	}

	cmd.Stderr = os.Stderr 
	cmd.Stdin = os.Stdin   

	if err := cmd.Run(); err != nil {
		fmt.Printf("whussh: %v\n", err)
	}
}

func executePipeline(commands [][]string) {
	var cmds []*exec.Cmd
	var prev *exec.Cmd

	for _, args := range commands {
		cmd := exec.Command(args[0], args[1:]...)
		if prev != nil {
			pipe, _ := prev.StdoutPipe()
			cmd.Stdin = pipe
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmds = append(cmds, cmd)
		prev = cmd
	}

	for _, cmd := range cmds {
		if err := cmd.Start(); err != nil {
			fmt.Printf("whussh: %v\n", err)
			return
		}
	}

	for _, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			fmt.Printf("whussh: %v\n", err)
		}
	}
}

func handleBuiltin(args []string) bool {
	if len(args) == 0 {
		return false
	}
	handler, exists := builtins.Builtins[args[0]]
	if !exists {
		return false
	}
	handler(args)
	return true
}