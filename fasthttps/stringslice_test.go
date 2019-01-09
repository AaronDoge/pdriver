package fasthttps

import "testing"

func TestSuccessStringInSlice(t *testing.T) {
	if ok := StringInSlice("a", []string{"a", "b"}); ok {
		t.Log("test PASS")
	} else {
		t.Error("test FAILED")
	}
}

func TestFailStringInSlice(t *testing.T) {
	if ok := StringInSlice("c", []string{"a", "b"}); ok {
		t.Error("test FAILED")
	} else {
		t.Log("test PASS")
	}
}

func TestFibonacci(t *testing.T) {
	t.Log(Fibonacci(0))
	t.Log(Fibonacci(1))
	t.Log(Fibonacci(2))
	t.Log(Fibonacci(5))
	t.Log(Fibonacci(6))
	t.Log(Fibonacci(10))
	t.Log(Fibonacci(20))
	t.Log(Fibonacci(64))
}
