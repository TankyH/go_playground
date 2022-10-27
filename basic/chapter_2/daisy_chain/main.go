package main

import "fmt"

// 值从通道 right 传给了 left, 并累加 1
func f(left, right chan int) {
	left <- 1 + <-right
}

// 信道的信道, 也看出来 goroutine 的轻量, 10w个线程之间的通信很快就跑完
func main() {
	const n = 100000

	// init channels
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	// 一个个管道拼接成一个长串的, 等待最右边的管道发送信号
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	// 最右边的管道发送信号, 值经过一个个管道去到最左边
	go func(c chan int) { c <- 1 }(right)

	// 输出累加后的 100001
	fmt.Println(<-leftmost)
}
