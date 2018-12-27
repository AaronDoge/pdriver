package main

import (
	"errors"
	"fmt"
)

var ErrDidNotWork = errors.New("Did not work!")

func DoTheThings(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing() 	// 我原来以为这里err会报错：上面已经定义过err。然而并没有。。。
										// 在一个{}作用域内，其外的变量可以被覆盖（使用:=），若不覆盖则默认是{}前最近的那个变量
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func DoTheThing(reallyDoIt bool) (err error) {
	var result string
	if reallyDoIt {
		result, err = tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}

func main() {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
}
