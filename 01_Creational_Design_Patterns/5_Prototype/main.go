package main

import "fmt"

// Prototype 是原型接口，定义了复制方法
type Prototype interface {
	Clone() Prototype
}

// Sheep 是具体原型类
type Sheep struct {
	Name   string
	Color  string
	Weight int
}

func (s *Sheep) Clone() Prototype {
	return &Sheep{
		Name:   s.Name,
		Color:  s.Color,
		Weight: s.Weight,
	}
}

func main() {
	originalSheep := &Sheep{
		Name:   "Dolly",
		Color:  "White",
		Weight: 50,
	}

	fmt.Println("Original sheep:", originalSheep)

	// 复制对象
	clonedSheep := originalSheep.Clone().(*Sheep)
	clonedSheep.Name = "Molly"
	clonedSheep.Weight = 55

	fmt.Println("Cloned sheep:", clonedSheep)
}

/*
在 clonedSheep := originalSheep.Clone().(*Sheep) 这行代码中，originalSheep.Clone()
 返回的是一个实现了 Prototype 接口的对象，
但是接口类型是不确定的。因此，我们需要使用类型断言 (*Sheep) 将
其转换为具体的 *Sheep 类型，以便我们可以修改该对象的属性。

originalSheep.Clone() 方法返回的是一个 Prototype 接口类型，
因此我们需要使用 (*Sheep) 来将其转换为 *Sheep 类型。
这样，在后续的代码中，我们就可以像处理 *Sheep 类型对象一样，对 clonedSheep 进行属性的修改。
*/
