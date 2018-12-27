package main

import "fmt"

/*
闭包实现一个计数器

 */
func counter() func() int {
	i := 0
	return func() int {
		// fmt.Println("value of i is ", i)
		i += 1
		return i
	}
}

func main() {
	c1 := counter()
	c2 := counter()

	fmt.Println("c() val is ", c1())
	fmt.Println("c() val is ", c1())
	fmt.Println("c() val is ", c1())
	fmt.Println("c() val is ", c2())
	fmt.Println("c() val is ", c2())
	fmt.Println("c() val is ", c1())

	/*
	输出：
	c() val is  1
	c() val is  2
	c() val is  3
	c() val is  1
	c() val is  2
	c() val is  4
	 */

}

