package main

/*
#include "utils.c"
 */
import "C"
import "fmt"

func main() {
	fmt.Println(C.add(2, 1))
	C.pri()
}