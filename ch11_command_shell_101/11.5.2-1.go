package main

import (
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile(`\s+`)
	tokens := r.Split(`echo hello`, -1)
	for i, t := range tokens {
		fmt.Println(i, t)
	}
}
