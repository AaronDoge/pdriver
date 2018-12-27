package main

import (
	"fmt"
	"pdriver/cmd"
	"time"
)

func main() {
	st := time.Now()

	cmd.Execute()

	//time.Sleep(1 * time.Second)
	et := time.Now()

	fmt.Println("start time", st)
	fmt.Println("end time", et)

	fmt.Println("delta time", et.Sub(st))
}
