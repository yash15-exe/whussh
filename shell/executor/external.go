package executor

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecutePipeline(commands [][]string) error {
	var cmds []*exec.Cmd
	var prev *exec.Cmd

	for _, args := range commands {
		cmd := exec.Command(args[0], args[1:]...)
		if prev != nil {
			// Connect previous command's stdout to current stdin
			r, _ := prev.StdoutPipe()
			cmd.Stdin = r
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmds = append(cmds, cmd)
		prev = cmd
	}

	// Start all commands
	for _, cmd := range cmds {
		if err := cmd.Start(); err != nil {
			return fmt.Errorf("failed to start command: %v", err)
		}
	}

	// Wait for all commands
	for _, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			return fmt.Errorf("command failed: %v", err)
		}
	}
	return nil
}