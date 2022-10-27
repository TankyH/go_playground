package concurrency

import (
	"sync"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second * 1)
	// 不是线程安全的
	t.Log(counter)
}

func TestGoroutine2(t *testing.T) {
	counter := 0
	for ; counter < 5000; counter++ {
		go func(i int) {
			i++
		}(counter)
	}
	time.Sleep(time.Second * 1)
	t.Log(counter)
}

func TestMutex(t *testing.T) {
	counter := 0
	mux := sync.Mutex{}
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() { mux.Unlock() }()
			mux.Lock()
			counter++
		}()
	}
	// 为了让输出可以先出来, 这里sleep 一下
	time.Sleep(time.Second * 1)
	t.Log(counter)
}

func TestWaitGroupWithMutex(t *testing.T) {
	counter := 0
	mux := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mux.Unlock()
				wg.Done()
			}()
			mux.Lock()
			counter++
		}()
	}
	wg.Wait()
	// 用waitgroup, 可以不用sleep了
	t.Log(counter)
}

//func TestWaitGroupWithChan(t *testing.T) {
//	counter := 0
//	mux := sync.Mutex{}
//	wg := sync.WaitGroup{}
//	for i := 0; i < 5000; i++ {
//		wg.Add(1)
//		go func() {
//			defer func() {
//				mux.Unlock()
//				wg.Done()
//			}()
//			mux.Lock()
//			counter++
//		}()
//	}
//	wg.Wait()
//	// 用waitgroup, 可以不用sleep了
//	t.Log(counter)
//}
