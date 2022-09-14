package basics

import (
	"math/rand"
	"time"
)

func RandomInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	// half-open interval [0,n)
	return rand.Intn(n)
}

// RandomType
// interface{} 是接口, 能表达所有类型, 在 go 1.18 用 any 代替 interface{}
// 类似 js object
func RandomType() any {
	r := RandomInt(5)
	m := map[int]any{
		1: 1,
		2: "string",
		3: 0.01,
		4: func() {},
	}

	return m[r]
}
