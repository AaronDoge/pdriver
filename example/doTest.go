package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type A struct {
	Name 	string
	Age 	int
}

func main() {

	output, _ := exec.Command("pgrep", "-f", "sdk").Output()
	pid := strings.TrimSpace(string(output))
	fmt.Println("pid is", pid)


}

func IsRunning() string {
	var pid string
	output, _ := exec.Command("pgrep", "-f", "sdk-gateway").Output()
	pid = strings.TrimSpace(string(output))

	return pid
}


func test() ([]A, *[]A) {
	var a  = A{
		Name: "ni",
		Age: 10,
	}

	var b = A{
		Name: "hao",
		Age: 11,
	}

	s := []A{a,b}

	return s, &s
}
