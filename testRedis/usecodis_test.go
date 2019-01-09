package testRedis

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	if val := Get("say"); val == "hello" {
		fmt.Println("value: ", val)
		fmt.Println("pass")
	} else {
		fmt.Println("Fail")
	}
}

func TestSet(t *testing.T) {
	if err := Set("nihao", "bonjour"); err != nil {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
