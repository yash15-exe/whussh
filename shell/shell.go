package shell

import (
	"bufio"
	"fmt"
	"os"
	"whussh/shell/parser"
	"whussh/shell/executor"
)

func Start(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Whussh shell!")
	for{
		fmt.Print("whussh>")
		if !scanner.Scan(){
			break
		}

		input:=scanner.Text()
		args:= parser.Parse(input)
		executor.Execute(args)
	}
}