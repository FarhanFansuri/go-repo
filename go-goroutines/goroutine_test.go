package goro

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	RunHelloWorld()
	fmt.Println("Loh!")

	time.Sleep(1 * time.Second)
}
