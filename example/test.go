package main

import "fmt"

func main() {
	var a = make(map[string]string)

	a["ni"] = "you"

	if val, ok := a["ni"]; ok {
		fmt.Println(val)
	}
}

//func main() {
//	//var mutex *sync.RWMutex
//	mutex := new(sync.RWMutex)
//
//	fmt.Println("Locking the lock...")
//	mutex.Lock()
//	fmt.Println("The lock has been locked...")
//
//	channels := make([]chan int, 4)
//	for i := 0; i < 4; i++ {
//		channels[i] = make(chan int)
//		go func(i int, c chan int) {
//			fmt.Println("locking... ", i)
//			mutex.Lock()
//			fmt.Println("locked ", i)
//			fmt.Println("unlocking... ", i)
//			mutex.Unlock()
//			c <- i
//		}(i, channels[i])
//	}
//	time.Sleep(time.Second * 2)
//	fmt.Println("Unlocking the lock...")
//	mutex.Unlock()
//	time.Sleep(time.Second)
//
//	for _, c := range channels {
//		<- c
//	}
//}
