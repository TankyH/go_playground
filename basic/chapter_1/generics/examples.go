package generics

import "fmt"

type MyType struct {
	Name string
	Type int
}

//TypeSwitch 有泛型之前我们用 type switch 来根据类型做不同的逻辑处理
func TypeSwitch(i interface{}) {
	//type switch
	switch i.(type) {
	case int:
		fmt.Println("int.")
	case string:
		fmt.Println("string.")
	case MyType:
		fmt.Println("my type.")
	default:
		fmt.Println("unknown type")
	}
}

func TypeGeneric[T any](i T) {
	// // 泛型不能使用 type switch, (不能 type assertion)
	//switch i.(type) {
	//case int:
	//	fmt.Println("int.")
	//case string:
	//	fmt.Println("string.")
	//case MyType:
	//	fmt.Println("my type.")
	//default:
	//	fmt.Println("unknown type")
	//}
}

type MyInt int

type MyGenericType interface {
	~int | float32 | float64 | string
	// ~ 是泛指 int 的类型, 比如我们之前定义的 MyInt
}

// Compare 使用 comparable 关键字进行比较,
// 但 comparable 只支持 ==, !=, 不能 >, <, >=, <=
// 需要用到自定义的类型来比较
//func Compare[T comparable](a, b T) T {
func Compare[T MyGenericType](a, b T) T {
	//	if b == a {
	if b < a {
		return b
	}
	return a
}

// type paramters
func MinGenerate[T MyGenericType](a, b T) T {
	return Compare(a, b)
}
