package main

import "fmt"

var globalVal = 2
var globalVal2 = 2

func check(i int) func() int {

	tmp := func() int {
		i ++
		fmt.Println("global value in inner func of check: ", globalVal)
		return i + globalVal
	}

	i += 2 	// tmp是一个引用了check的局部变量i的闭包，程序执行check，会将闭包以及其引用的变量一起放到内存中
			// 但并不执行闭包的内容，直到调用闭包函数的时候才会执行
			// 闭包引用的变量的值，是该变量(包括全局和局部变量)在闭包函数被调用前最后的值，注意是调用前，而不是创建前。
	// fmt.Println("global in check: ", globalVal)
	return tmp
}

func check_2(i int) func() int {

	tmp2 := func(n int) func() int{
		tmp := func() int {
			n ++
			return n + globalVal2
		}
		return tmp
	} (i)

	i += 2 	// 此处的i的值不会影响闭包的返回值了，tmp引用了匿名函数tmp2的参数n，而i以 值传递 的方式赋值给n的
			// 所以闭包tmp中的返回值不会受到i的值影响。
	return tmp2 // tmp2是匿名函数，返回一个闭包
}

func main() {
	ch1 := check(0)
	// globalVal = 3  // 放在这里，在ch1()，ch1_1()被调用前，因此ch1()和ch1_1()中globalVal的值是一样的，所以返回值一样，都是6
	ch1_1 := check(0)
	//ch2 := check(1) // 新建一个闭包，（function factory）

	fmt.Println("ch1: ", ch1()) // ch1: 5
	globalVal = 3	// ch1()被调用后，ch1_1()被调用前globalVal最后的值被修改，ch1()中和ch1_1()中globalVal的值分别为2和3
					// 因此ch1()和ch1_1()的返回值不一样，分别为5和6

	fmt.Println("ch1_1: ", ch1_1()) // ch1_1: 6
	//fmt.Println(ch2())

	fmt.Println("------------------")

	ch2 := check_2(0)

	ch2_1 := check_2(1)

	fmt.Println("ch2: ", ch2()) // ch2:  3

	globalVal2 = 3 	// globalVal2的值还是会影响闭包的值的
	fmt.Println("ch2_1: ", ch2_1()) // ch2_1:  5

}

func init() {

}
