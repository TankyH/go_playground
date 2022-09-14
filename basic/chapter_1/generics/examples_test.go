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
	var a int
	var b = 2
	var c float32
	t.Log(Compare(a, b))
	//t.Log(Compare(a, c))
	t.Log(a, b, c)
}
