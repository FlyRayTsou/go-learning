package main

import (
	"os"
	"os/exec"
	"runtime"
)

func main() {
	var cmd *exec.Cmd
	println(runtime.GOOS)
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/C", "timeout 5")

	} else {
		println(os.Getenv("SHELL"))
		cmd = exec.Command(os.Getenv("SHELL"), "-c", "sleep 5")
	}
	cmd.Run()
}
