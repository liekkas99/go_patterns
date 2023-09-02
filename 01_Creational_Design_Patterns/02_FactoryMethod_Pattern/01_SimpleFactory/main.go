package main

import (
	"fmt"
)

// Product 是一个接口，代表产品的抽象
type Product interface {
	ShowInfo()
}

// ConcreteProductA 是一个具体产品，实现了 Product 接口
type ConcreteProductA struct{}

func (p ConcreteProductA) ShowInfo() {
	fmt.Println("ConcreteProductA")
}

// ConcreteProductB 是另一个具体产品，也实现了 Product 接口
type ConcreteProductB struct{}

func (p ConcreteProductB) ShowInfo() {
	fmt.Println("ConcreteProductB")
}

// Factory 是工厂类，负责创建产品
type Factory struct{}

func (f Factory) CreateProduct(productType string) Product {
	switch productType {
	case "A":
		return ConcreteProductA{}
	case "B":
		return ConcreteProductB{}
	default:
		return nil
	}
}

func main() {
	factory := Factory{}

	productA := factory.CreateProduct("A")
	productB := factory.CreateProduct("B")

	productA.ShowInfo() // 输出: ConcreteProductA
	productB.ShowInfo() // 输出: ConcreteProductB
}
