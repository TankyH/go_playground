package concurrency

import "sync"

type Singleton struct {
}

var singleton Singleton
var once sync.Once

func Run() {
	once.Do(func() {
		singleton = Singleton{}
	})
}
