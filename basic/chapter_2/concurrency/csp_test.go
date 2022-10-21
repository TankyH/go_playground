package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func Producer(i int) chan int {
	c := make(chan int)
	c <- i
	return c
}

func Consumer(c chan int) {
	i := <-c
	fmt.Println(i)
}

func TestChannel(t *testing.T) {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	for {
		select {
		case i := <-c:
			fmt.Println(i)
		case <-time.After(time.Second * 2):
			fmt.Println("time out!!")

		}
	}

}
