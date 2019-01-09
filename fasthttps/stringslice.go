package fasthttps

import "fmt"

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Fibonacci(n int) int {
	var fibseq = make([]int, n + 1)

	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	fibseq[0] = 0
	fibseq[1] = 1

	for i := 2; i < n + 1; i++ {
		fibseq[i] = fibseq[i - 1] + fibseq[i - 2]
	}

	fmt.Println(fibseq)

	return fibseq[n]
}
