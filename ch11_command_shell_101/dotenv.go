package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	filename := flag.String("e", ".env", ".env file name to read")
	// println(filename) 0xc0000960a0
	flag.Parse()
	cmdName := flag.Arg(0)
	// println(cmdName) ./my-printenv.bash
	args := flag.Args()[1:]
	// println(args) [0/0]0xc0000b2010
	flag.Args()

	cmd := exec.Command(cmdName, args...)

	envs := os.Environ()
	dotenvs, _ := godotenv.Read(*filename)
	for key, value := range dotenvs {
		envs = append(envs, key+"="+value)
	}
	cmd.Env = envs
	o, err := cmd.CombinedOutput()
	fmt.Println(string(o))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
