package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2021, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	parse, _ := time.Parse(time.RFC3339, "2021-11-24T15:04:05Z")
	fmt.Println(parse)
}
