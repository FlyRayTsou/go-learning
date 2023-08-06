package main

import (
	"fmt"
)

func main() {
	var cmd string
	/*
	 * The fmt.Scanf() function in Go language scans the input texts which is given in the standard input,
	 * reads from there and stores the successive space-separated values into successive arguments as determined by the format.
	 */
	fmt.Scanf("%s", &cmd)
	fmt.Printf("You input cmd : %s", cmd)
}
