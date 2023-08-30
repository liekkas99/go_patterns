package main

import "fmt"

// AbstractProductA 是抽象产品A接口
type AbstractProductA interface {
	DisplayInfo()
}

// ConcreteProductA1 是具体产品A1
type ConcreteProductA1 struct{}

func (p ConcreteProductA1) DisplayInfo() {
	fmt.Println("ConcreteProductA1")
}

// ConcreteProductA2 是具体产品A2
type ConcreteProductA2 struct{}

func (p ConcreteProductA2) DisplayInfo() {
	fmt.Println("ConcreteProductA2")
}

// AbstractProductB 是抽象产品B接口
type AbstractProductB interface {
	DisplayInfo()
}

// ConcreteProductB1 是具体产品B1
type ConcreteProductB1 struct{}

func (p ConcreteProductB1) DisplayInfo() {
	fmt.Println("ConcreteProductB1")
}

// ConcreteProductB2 是具体产品B2
type ConcreteProductB2 struct{}

func (p ConcreteProductB2) DisplayInfo() {
	fmt.Println("ConcreteProductB2")
}

// AbstractFactory 是抽象工厂接口
type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

// ConcreteFactory1 是具体工厂1，用于创建具体产品簇1
type ConcreteFactory1 struct{}

func (f ConcreteFactory1) CreateProductA() AbstractProductA {
	return ConcreteProductA1{}
}

func (f ConcreteFactory1) CreateProductB() AbstractProductB {
	return ConcreteProductB1{}
}

// ConcreteFactory2 是具体工厂2，用于创建具体产品簇2
type ConcreteFactory2 struct{}

func (f ConcreteFactory2) CreateProductA() AbstractProductA {
	return ConcreteProductA2{}
}

func (f ConcreteFactory2) CreateProductB() AbstractProductB {
	return ConcreteProductB2{}
}

func main() {
	factory1 := ConcreteFactory1{}
	factory2 := ConcreteFactory2{}

	productA1 := factory1.CreateProductA()
	productB1 := factory1.CreateProductB()

	productA2 := factory2.CreateProductA()
	productB2 := factory2.CreateProductB()

	productA1.DisplayInfo() // 输出: ConcreteProductA1
	productB1.DisplayInfo() // 输出: ConcreteProductB1

	productA2.DisplayInfo() // 输出: ConcreteProductA2
	productB2.DisplayInfo() // 输出: ConcreteProductB2
}
