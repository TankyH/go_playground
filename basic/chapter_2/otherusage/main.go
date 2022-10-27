package main

import (
	"fmt"
	"time"
)

func main() {
	chanSample1()
	//chanSample2()
}

func chanSample1() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		ch1 <- 3
		close(ch1)
	}()

	go func() {
		time.Sleep(time.Second * 6)
		ch2 <- 6
		close(ch2)
	}()

	var ok1, ok2 bool
	for {
		select {
		case x := <-ch1:
			ok1 = true
			fmt.Println(x)
		case x := <-ch2:
			ok2 = true
			fmt.Println(x)
		}

		if ok1 && ok2 {
			break
		}
	}

	fmt.Println("program end")
}

func chanSample2() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second * 3)
		ch1 <- 3
		close(ch1)
	}()

	go func() {
		time.Sleep(time.Second * 6)
		ch2 <- 6
		close(ch2)
	}()

	for {
		select {
		case x, ok := <-ch1:
			if !ok {
				ch1 = nil // 设置为 nil, 不再排出
			} else {
				fmt.Println(x)
			}
		case x, ok := <-ch2:
			if !ok {
				ch2 = nil
			} else {
				fmt.Println(x)
			}
		}
		if ch1 == nil && ch2 == nil {
			break
		}
	}
	fmt.Println("program end")
}
