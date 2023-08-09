package main

import (
	"errors"
	"io"
	"log"

	"github.com/peterh/liner"
)

func main() {

	line := liner.NewLiner()
	line.SetCtrlCAborts(true)
	// var (
	// 	names = []string{"john", "james", "mary", "nancy"}
	// )
	// line.SetCompleter(func(line string) (c []string) {
	// 	for _, n := range names {
	// 		if strings.HasPrefix(n, strings.ToLower(line)) {
	// 			c = append(c, n)
	// 		}
	// 	}
	// 	return
	// })
	for {
		if cmd, err := line.Prompt(" "); err == nil {
			if cmd == "" {
				continue
			}
			// ここでコマンドを処理する
			log.Print(cmd)
		} else if errors.Is(err, io.EOF) {
			break
		} else if err == liner.ErrPromptAborted {
			log.Print("Aborted")
			break
		} else {
			log.Print("Error reading line: ", err)
		}
	}
}
