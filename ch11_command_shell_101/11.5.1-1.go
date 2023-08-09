package main

import (
	"fmt"
)

func main() {
	var cmd, cmd2 string
	// Calling Scanf() function for
	// scanning and reading the input
	// texts given in standard input

	// rayt@RaydeMacBook-Pro ch11_command_shell_101 % go run 11.5.1-1.go
	// cd ..
	fmt.Scanf("%s %s", &cmd, &cmd2)

	// You input cmd : cd .. .
	fmt.Printf("You input cmd : %s %s .\n", cmd, cmd2)
}
