package generics

import "testing"

func TestTypeSwitch(t *testing.T) {
	var a int
	TypeSwitch(a)

	var s string
	TypeSwitch(s)

	var mt MyType
	TypeSwitch(mt)

	var f float32
	TypeSwitch(f)
}

func TestTypeGeneric(t *testing.T) {
	// type switch 不能在 泛型约束 中使用.

	// cannot use type switch on type parameter value
	// i (variable of type T constrained by any)
	var a int
	TypeGeneric(a)

	var s string
	TypeGeneric(s)

	var mt MyType
	TypeGeneric(mt)

	var f float32
	TypeGeneric(f)
}

func TestCompare(t *testing.T) {
	a, b := 1, 2
	mi := MyInt(1)
	mi2 := MyInt(2)
	var c float32
	var d float64

	t.Log(MinGenerate(a, b))
	t.Log(MinGenerate(mi, mi2))
	//t.Log(MinGenerate(a, mi)) // ~int 类型不同也是不能直接比较
	//t.Log(Compare(c, d)) // 同理, float32, float64 不能比较

	t.Log("-------")
	t.Log(a, b, c, d)
}
