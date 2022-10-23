package goroutines

import (
	"fmt"
	"sync"
	"testing"
	_ "time"
)

func TestPool(t *testing.T) {
	group := sync.WaitGroup{}
	pool := sync.Pool{
		New: func() interface{} {
			return "Vanhautten"
		},
	}

	pool.Put("Azie")
	pool.Put("Melza")
	pool.Put("Pratama")

	for i := 0; i < 10; i++ {
		go func() {
			defer group.Done()
			group.Add(1)

			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	group.Wait()
}
