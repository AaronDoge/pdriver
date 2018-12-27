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
