package goroutines

import (
	"fmt"
	"testing"
	"time"
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

	time.Sleep(1 * time.Second)
}

func GiveMeResponse(channel chan string) {
	channel <- "Azie Melza Pratama"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(1 * time.Second)
}

func OnlyIn(c chan<- string) {
	c <- "Azie Melza Pratama"
}

func OnlyOut(c <-chan string) {
	data := <-c
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(1 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Azie"
		channel <- "Melza"
		channel <- "Pratama"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Selesai")
}
