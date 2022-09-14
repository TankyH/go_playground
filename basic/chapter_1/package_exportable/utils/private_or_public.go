package utils

import "fmt"

// myInt 小写 代表 protected, 包外不可以访问
type myInt int

// getMyInt
func getMyInt(m myInt) int {
	fmt.Println("can not be exposed outside package.")
	return int(m)
}

// MyInt 大写 代表 public, 包外可以访问
type MyInt int

func GetMyInt(m MyInt) int {
	fmt.Println("can be exposed outside package.")
	return int(m)
}
