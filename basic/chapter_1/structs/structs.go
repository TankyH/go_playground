package structs

import "fmt"

type Creature interface {
	Run()
}

// Animal 仅仅是实现了 Run(), 实现了 Creature 接口
type Animal struct {
	name string
}

func NewAnimal(name string) *Animal {
	return &Animal{name: name}
}

func (a *Animal) Run() {
	fmt.Printf("%s Running...\n", a.name)
}

// Dog "继承"了 Animal,
// 即使它没有显式的实现 Run(), 因为 Animal 带了 Run(),
// 所以 Dog 也是 Creature 的一个实现
type Dog struct {
	*Animal
	//name string
	dogCollar string
}

//// Name 还是 Animal 的, 并不能直接使用 name
//func NewDog(name string) *Dog {
//	return &Dog{Name: name, dogCollar: name}
//}

func NewDog(name string) *Dog {
	animal := &Animal{name: name} // 与其说是继承, 不如说是组合, mixin
	return &Dog{Animal: animal, dogCollar: name}
}

func (d *Dog) Swim() {
	fmt.Printf("%s swimming...\n", d.name)
}

// SetDogCollar setter function ( getter 一般不另外写方法, 就直接点出属性来
func (d *Dog) SetDogCollar(collar string) {
	d.dogCollar = collar
}

// Duck "继承"了 Animal
type Duck struct {
	*Animal
	egg string
}

func NewDuck(name string) Duck {
	return Duck{Animal: &Animal{name: name}, egg: name}
}

// Fly duck 特有的方法
func (d Duck) Fly() {
	fmt.Printf("%s flying...\n", d.name)
}

// Hatch duck 特有的方法
func (d Duck) Hatch() string {
	return fmt.Sprintf("little %s", d.name)
}

// 值传递 or 指针传递

func ValueReceiver(duck Duck) {
	duck.egg = "chick"
	fmt.Printf("pointer receiver: egg: %s\n", duck.egg)
}

func PointerReceiver(duck *Duck) {
	duck.name = "Kitty"
	fmt.Printf("value receiver: duck name %s\n", duck.name)
}
