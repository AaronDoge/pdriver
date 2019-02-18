package panicrecovery

import (
	"fmt"
	"time"
)

func PanicRecoveryGoroutine(user string) {

	go func() {
		defer func() {
			fmt.Println("defer 1st...")
			if err := recover(); err != nil {
				fmt.Println("1st panic message:", err)
			}
		}()
		defer func() {
			fmt.Println("defer here....")
			if err := recover(); err != nil {
				fmt.Println("panic message: ", err)
				fmt.Println("recover success...")
			}
		}()

		if user == "" {
			panic("user is nill!!!")
		}

		panic("2nd panic")

		fmt.Println("after panic in a goroutine")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("function end...")

}

func PanicRecovery(user string) string {
	defer func() {
		fmt.Println("defer here....")
		if err := recover(); err != nil {
			fmt.Println("panic message: ", err)
			fmt.Println("recover success...")
		}
	}()

	if user == "" {
		panic("user is nill!!!")
	}

	name := user + "aaron"
	return name
}