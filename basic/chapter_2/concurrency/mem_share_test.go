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

func TestGoroutine3(t *testing.T) {
	counter := 0
	mux := sync.Mutex{}
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() { mux.Unlock() }()
			mux.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second * 1)
	t.Log(counter)
}

func TestGoroutine4(t *testing.T) {
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
	t.Log(counter)
}
