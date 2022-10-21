package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoRoutines(t *testing.T) {
	go RunHelloWorld()

	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}
