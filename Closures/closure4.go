package main

import "fmt"

/*
使用闭包模拟面向对象
实现栈push和pop
*/

func NewStack() func(action string) func(dt ...string) string {
	data := make([]string, 5)
	last := -1

	push := func(e ...string) string {
		last += 1
		data[last] = e[0]
		return ""
	}

	pop := func(e ...string) string {
		if last == -1 {
			return ""
		}
		last = last - 1
		return data[last+1]
	}

	return func(action string) func(dt ...string) string {
		if action == "push" {
			return push
		} else if action == "pop" {
			return pop
		}
		return nil
	}
}

func main() {
	stack := NewStack()
	stack("push")("test01")
	stack("push")("test02")
	stack("push")("test03")
	stack("push")("test04")
	fmt.Println("pop is ", stack("pop")())
	fmt.Println("pop is ", stack("pop")())
	fmt.Println("pop is ", stack("pop")())
}

/*
输出：
pop is  test04
pop is  test03
pop is  test02
 */
