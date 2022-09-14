package structs

import "fmt"

type Creature interface {
	Run()
}

// Animal
type Animal struct {
	name string
}

func NewAnimal(name string) *Animal {
	return &Animal{name: name}
}

func (a *Animal) Run() {
	fmt.Printf("%s Running...\n", a.name)
}

// Dog
type Dog struct {
	*Animal
	//name string
	dogCollar string
}

// Name 还是 Animal 的
//func NewDog(name string) *Dog {
//	return &Dog{Name: name, dogCollar: name}
//}

func NewDog(name string) *Dog {
	animal := &Animal{name: name}
	return &Dog{Animal: animal, dogCollar: name}
}

func (d *Dog) Swim() {
	fmt.Printf("%s swimming...\n", d.name)
}

// SetDogCollar setter function
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

func (d Duck) Fly() {
	fmt.Printf("%s flying...\n", d.name)
}

func (d Duck) Hatch() string {
	return "little" + d.egg
}

func ValueReceiver(duck Duck) {
	duck.egg = "chick"
	fmt.Printf("pointer receiver: egg: %s\n", duck.egg)
}

func PointerReceiver(duck *Duck) {
	duck.name = "Kitty"
	fmt.Printf("value receiver: duck name %s\n", duck.name)
}
