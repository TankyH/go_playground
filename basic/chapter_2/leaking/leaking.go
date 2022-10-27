package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func networkCall() int {
	// just mocking
	time.Sleep(1 * time.Second)
	return 1
}

func forgottenSender(ch chan int) {
	data := networkCall()

	ch <- data
}

func handler() error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	ch := make(chan int)
	go forgottenSender(ch)

	select {
	case data := <-ch:
		{
			fmt.Printf("Received data! %s", data)

			return nil
		}

	case <-ctx.Done():
		{
			return errors.New("timeout: handler cancelled! I'm out")
		}
	}
}

func main() {
	err := handler()
	fmt.Println(err)
}

func handlerAfterError() error {
	ch := make(chan int)
	go forgottenSender(ch)

	err := otherLogic()
	if err != nil {
		// 当 error 出现会提早退出程序, 下面的 chan 接受器
		return errors.New("data is invalid")
	}
	data := <-ch
	fmt.Println(data)
	return nil
}

func otherLogic() error {
	// just mock here
	return errors.New("something mistake here")
}
