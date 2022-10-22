package goroutines

import (
	"fmt"
	"testing"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		channel <- "Azie Melza Pratama"
		fmt.Println("Selesai Mengirim Data  ke Channel")
	}()

	data := <-channel
	fmt.Println(data)
}
