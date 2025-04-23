package utils

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func TrapInterrupt() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		fmt.Println("\nwhussh: type 'exit' to quit")
	}()
}