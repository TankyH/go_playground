package structs

import (
	"fmt"
	"testing"
)

func TestAnimal(t *testing.T) {
	var animal Creature
	animal = NewAnimal("MR. Nobody")
	animal.Run()

	//var duck Creature
	// init
	duck := Duck{
		Animal: &Animal{
			name: "duck",
		},
		egg: "egg",
	}
	duck.Run()
	duck.Fly()
	egg := duck.Hatch()
	fmt.Println("duck egg:", egg)

	// 习惯上我们会加个 NewXxx 的工厂方法来实例化
	dog := NewDog("Dog")
	dog.Run()
	dog.Swim()
	dog.SetDogCollar("Husky")
	fmt.Println(dog.dogCollar)

}

func TestReceiver(t *testing.T) {
	// 值接收器, 指针接收器
	d1 := NewDuck("Daisy")
	d2 := NewDuck("Donald")

	fmt.Println(d1.egg)
	ValueReceiver(d1)
	fmt.Println(d1.egg) // egg 变了没?

	fmt.Println("------")

	fmt.Println(d2.name)
	PointerReceiver(&d2)
	fmt.Println(d2.name) // name 变了没?

	// 指针接受器
	// 不会copy值, 能维护属性, 不是并发安全的

	// 值接受器
	// 不需要维护属性的场景, 并发安全, 语义明确, 对于大struct会有copy开销

	// 对于小struct, 和基本类型, GC 开销不大, 这时值/指针接收器差别不大
}
