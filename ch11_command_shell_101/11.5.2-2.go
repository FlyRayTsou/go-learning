package main

import (
	"fmt"
	"strings"

	"github.com/google/shlex"
)

func parseCmd(cmdStr string) (cmd string, args []string, err error) {
	l := shlex.NewLexer(strings.NewReader(cmdStr))
	// cmd: echo
	cmd, err = l.Next()
	if err != nil {
		return
	}
	for {
		// token: hello world
		token, err := l.Next()
		if err != nil {
			break
		}
		args = append(args, token)
	}
	return
}

func main() {
	cmdStr := "echo 'hello world' || echo 'happy world'"
	// cmdStr := "| ; || && < > >> 2> 2>> &> &>>"
	cmd, args, err := parseCmd(cmdStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Command is  \n", cmd)
	for i := 0; i < len(args); i++ {
		fmt.Println("%s\n", args[i])
	}
}
