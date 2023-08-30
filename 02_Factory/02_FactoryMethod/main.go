package main

import (
	"fmt"
)

// Product 是产品接口
type Product interface {
	ShowInfo()
}

// ConcreteProductA 是具体产品A
type ConcreteProductA struct{}

func (p ConcreteProductA) ShowInfo() {
	fmt.Println("ConcreteProductA")
}

// ConcreteProductB 是具体产品B
type ConcreteProductB struct{}

func (p ConcreteProductB) ShowInfo() {
	fmt.Println("ConcreteProductB")
}

// Factory 是工厂接口
type Factory interface {
	CreateProduct() Product
}

// ConcreteFactoryA 是具体工厂A，用于创建 ConcreteProductA
type ConcreteFactoryA struct{}

func (f ConcreteFactoryA) CreateProduct() Product {
	return ConcreteProductA{}
}

// ConcreteFactoryB 是具体工厂B，用于创建 ConcreteProductB
type ConcreteFactoryB struct{}

func (f ConcreteFactoryB) CreateProduct() Product {
	return ConcreteProductB{}
}

func main() {
	factoryA := ConcreteFactoryA{}
	factoryB := ConcreteFactoryB{}

	productA := factoryA.CreateProduct()
	productB := factoryB.CreateProduct()

	productA.ShowInfo() // 输出: ConcreteProductA
	productB.ShowInfo() // 输出: ConcreteProductB
}
