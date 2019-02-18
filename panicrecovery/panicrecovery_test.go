package panicrecovery

import (
	"fmt"
	"testing"
)

func TestPanicRecoveryGoroutine(t *testing.T) {
	PanicRecoveryGoroutine("")
}

func TestPanicRecovery(t *testing.T) {
	name := PanicRecovery("")
	fmt.Println("name is", name)
}