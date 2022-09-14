package basics

import (
	"fmt"
	"testing"
	"unsafe"
)

type myInt int

/*
 * 同目录下有文件 x.go, 可以在 x_test.go 写单测, 单测方法 Test 开头.
 * 顺便也看看 unit test 怎么 assert.
 * go build 不会编译 test, 不需要担心测试写多了会影响编译速度
 */
func TestTypes(t *testing.T) {

	// int
	var a int
	b := 0 // go 编译器自动类型推断, 不需要显式的指定 int
	t.Log(a, b)

	mi := myInt(1) // 自定义一个 "int" 类型
	t.Log(mi)
	//t.Log(a == mi) // 自定义的 MyInt 不能直接和 int 比较

	// float
	var (
		//f1 float32
		f1 float64
		f2 float64
	)
	f1 = 0.1
	f2 = .01

	t.Log(f1 == f2)

	// boolean
	c := true
	c = false
	d := a == b
	t.Logf("a:%T, b:%T, c:%T, d:%T", a, b, c, d)

	s := "abc"      // 编译器会自动类型推断, 初始化时可以省去类型的声明
	balance := 0.01 // float64
	t.Logf("s:%T, balance:%T", s, balance)

	// byte, rune, uintptr
	var (
		bt  byte
		r   rune
		ptr uintptr // 指针地址
	)
	bt = 0xf
	r = '中'
	t.Log(bt, r, ptr)
	t.Logf("r: %T, %v, %+v", r, r, r)
}

func TestArray(t *testing.T) {
	// 数组是定长的
	a := [3]int{1, 2, 3}   // 指定长度
	b := [...]int{1, 2, 3} // 懒得数了, ...自动定长度

	// forEach
	for i := range b {
		t.Log(i)
	}
	t.Log(a == b)   // 值的比较
	t.Log(&a == &b) // 指针地址比较
	t.Log(unsafe.Pointer(&a), unsafe.Pointer(&b))

	c := [...]int{1, 2, 3, 4} // 不同长度可以比较吗
	d := [...]myInt{1, 2, 3}  // 同时 ~int 能比较吗
	t.Log(c, d)

	//t.Log(a == d) // 自定义的 MyInt 也不能直接和 int 比较

	//t.Log(a == c)   // 长度不同的比较
}

// Slice 切片
func TestSlice(t *testing.T) {
	//slice len, cap
	a := []int{1, 2, 3}    // slice
	b := [...]int{1, 2, 3} // array

	// 观察 cap 的数值
	for i := 0; i < 10; i++ {
		a = append(a, i)
		t.Logf("a: %v, len: %v, cap: %v", a, len(a), cap(a))
	}
	t.Logf("------")

	//for i := 0; i < 10; i++ {
	//	// 数组是定长的
	//	b = append(b, i)
	//	t.Logf("b: %v, len: %v, cap: %v", b, len(b), cap(b))
	//}

	for i := range b {
		a = append(a, i)
		t.Logf("a: %v, len: %v, cap: %v", a, len(a), cap(a))
	}

	t.Logf("a type: %T, len: %v, cap: %v", a, len(a), cap(a))
	t.Logf("b type: %T", b)

	t.Log(&a, &b)
	//t.Log(a == b) // slice, array 不能直接比较
}

// Map
func TestMap(t *testing.T) {
	//m := new(map[string]int)  // new pointer
	m := make(map[string]int) // make map object
	m["a"] = 1
	fmt.Printf("m: %+v\n", m)
	fmt.Println("=====")

	m1 := map[string]int{
		"a": 1,
	}
	// check key exist
	if v, ok := m1["a"]; ok {
		fmt.Printf("err: %v, v: %v\n", ok, v)
	}
	// check key not exist
	if v, ok := m1["xxx"]; ok {
		fmt.Printf("err: %v, v: %v\n", ok, v)
	} else {
		fmt.Printf("not exists\n")
	}

	fmt.Println("=====")

	// iterate key, value of the map
	for k, v := range m1 {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
}

// Set golang 没有set
func TestSet(t *testing.T) {
	// go 没有 set 结构, 用 map 代替
	set := map[string]bool{
		"a": true,
		"b": false,
	}
	if v, ok := set["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("key not exists")
	}
}
