package basics

import (
	"fmt"
	"testing"
)

func TestCondition(t *testing.T) {
	var a, b = 1, 2
	if a == b {
		fmt.Printf("a: %d == b: %d\n", a, b)
	}

	s := ""
	//类型不同不能比较
	//if a == s {
	//	fmt.Printf("mismatched basics int and string")
	//}
	t.Log(s) // 当有变量定义了, 但没有被使用, 也会有编译问题, 比较严格
}

func TestSwitch1(t *testing.T) {
	a := RandomInt(5)

	switch a {
	case 1:
		// 不需要加 break
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	//case "":
	//	fmt.Println("类型需要对应")
	default:
		fmt.Println("a above 2")
	}
}

func TestSwitch2(t *testing.T) {
	a := RandomInt(5)
	b := RandomInt(5)
	// 条件可以不放在 switch, 对于没有互斥条件的场景
	switch {
	case a > b:
		fmt.Println("1")
	case a == b:
		fmt.Println("2")
	case a < b:
		fmt.Println("3")
	default:
		fmt.Println("other")
	}
}

func TestSwitch3(t *testing.T) {
	a := RandomType()
	// type switch 根据类型走不同逻辑的场景,
	// 对于 golang 这类强类型语言, 可以增加灵活性
	switch a.(type) {
	case int:
		// 不需要加 break
		fmt.Println("a is an int")
	case float64:
		fmt.Println("a is a float")
	case string:
		fmt.Println("a is a string")
	default:
		fmt.Println("type unknown")
	}
}

func TestLoop(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Printf("i: %v\n", i)
	}

	fmt.Println("------")

	i := 0
	for { // while
		if i >= 5 {
			break
		}
		fmt.Printf("i: %v\n", i)
		i++
	}

}
