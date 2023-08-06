package main

import (
	"io"
	"os/exec"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	c1 := exec.Command("echo 123") // Replace with your actual command
	c2 := exec.Command("echo 456") // Replace with your actual command

	reader, writer := io.Pipe()
	c1.Stdout = writer
	c2.Stdin = reader

	wg.Add(2)

	go func() {
		c1.Start()
		c1.Wait()
		writer.Close()
		wg.Done()
	}()

	go func() {
		c2.Start()
		c2.Wait()
		wg.Done()
	}()

	wg.Wait()
}
