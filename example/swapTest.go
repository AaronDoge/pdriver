package main

import "fmt"

func main() {
	var arr = []int{0,1,2,3,4,5}

	arr[0], arr[1] = arr[1], arr[0]
	//tmp := arr[0]
	//arr[0] = arr[1]
	//arr[1] = tmp

	fmt.Println(arr)	// [1 0 2 3 4 5]
}
