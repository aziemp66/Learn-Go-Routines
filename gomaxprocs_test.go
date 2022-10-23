package goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			time.Sleep(3 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total Cpu : ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread : ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Gourotine : ", totalGoroutine)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			time.Sleep(3 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total Cpu : ", totalCpu)

	runtime.GOMAXPROCS(12)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread : ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Gourotine : ", totalGoroutine)

	group.Wait()
}
