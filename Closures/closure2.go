package main

import "fmt"

func doTest(num int) (funcArr []func(name string) string) {

	for i := 0; i < num; i++ {
		funcArr = append(funcArr, func(name string) string {
			fmt.Printf("check value of %s: %d \n", name, i)
			return name
		})
	}
	return funcArr
}

func doTest2(num int) (funcArr []func(name string) string) {
	for i := 0; i < num; i++ {
		//tmp := func(n int) func(name string) string {
		//	return func(name string) string {
		//		fmt.Printf("check value of %s: %d \n", name, n)
		//		return name
		//	}
		//}(i) 	// i以值传递的方式传入
		//funcArr = append(funcArr, tmp)

		// 上面也可以如下这样写，更简练一点
		funcArr = append(funcArr, func(n int) func(name string) string {
			return func(name string) string {
				fmt.Printf("check value of %s: %d \n", name, n)
				return name
			}
		}(i))
	}
	return funcArr
}

func main() {
	checkTest := doTest(3)
	fmt.Println("name is ", checkTest[0]("test1"))	// 输出3，是for循环中i最后的值
	fmt.Println("name is ", checkTest[1]("test2"))
	fmt.Println("name is ", checkTest[2]("test3"))

	fmt.Println("------------------")

	checkTest2 := doTest2(3)
	fmt.Println("name is ", checkTest2[0]("test1"))
	fmt.Println("name is ", checkTest2[1]("test2"))
	fmt.Println("name is ", checkTest2[2]("test3"))
}

/*
输出：
	check value of test1: 3
	name is  test1
	check value of test2: 3
	name is  test2
	check value of test3: 3
	name is  test3
	------------------
	check value of test1: 0
	name is  test1
	check value of test2: 1
	name is  test2
	check value of test3: 2
	name is  test3
 */
