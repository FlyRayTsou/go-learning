package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t1 := time.Second
	fmt.Println("===== t1 =====")
	fmt.Println((t1.String()))
	t2 := 5 * time.Second
	fmt.Println("===== t2 =====")
	fmt.Println((t2.String()))
	t3 := 10 * time.Millisecond
	fmt.Println("===== t3 =====")
	fmt.Println((t3.String()))
	t4 := 10 * time.Nanosecond
	fmt.Println("===== t4 =====")
	fmt.Println((t4.String()))
	t5, t5error := time.ParseDuration("10m30s")
	fmt.Println("===== t5 =====")
	fmt.Println(t5)
	fmt.Println(t5error)
	t6, t6error := time.ParseDuration("10m3ks")
	fmt.Println("===== t6 =====")
	fmt.Println(t6)
	fmt.Println(t6error)

	now := time.Now()
	fmt.Println("===== now =====")
	fmt.Println(now)

	specificUTC := time.Date(2000, time.August, 26, 14, 45, 30, 0, time.UTC)
	fmt.Println("===== specificUTC =====")
	fmt.Println(specificUTC)

	specificLocal := time.Date(2000, time.August, 26, 14, 45, 30, 0, time.Local)
	fmt.Println("===== specificLocal =====")
	fmt.Println(specificLocal)

	timeParse, timeParseError := time.Parse(time.Kitchen, "11:30AM")
	fmt.Println("===== timeParse =====")
	fmt.Println(timeParse)
	fmt.Println(timeParseError)

	timeParse2, timeParseError2 := time.Parse(time.DateOnly, "2023-08-22")
	fmt.Println("===== timeParse2 =====")
	fmt.Println(timeParse2)
	fmt.Println(timeParseError2)

	epoch := time.Unix(1503673200, 0)
	fmt.Println("===== epoch =====")
	fmt.Println(epoch)

	// time.Time
	fmt.Println("===== Add.Hour =====")
	fmt.Println(time.Now().Add(3 * time.Hour))

	fmt.Println("===== Sub.Hour =====")
	fileInfo, _ := os.Stat("README.md")
	fmt.Println("%v前", time.Now().Sub(fileInfo.ModTime()))

	fmt.Println("===== Round.Hour =====")
	fmt.Println(time.Now().Round(time.Hour))
}
