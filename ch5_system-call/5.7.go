package main

import (
	"fmt"
	"log"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

func main() {
	var writer Writer
	count, err := writer.Write([]byte("hello"))
	if err != nil {
		log.Fatal(err.Error())
		fmt.Printf("count:%d", count)
	}
}
