package main

import "fmt"

// Implementor 是实现部分的接口，定义了支撑结构的方法
type Implementor interface {
	BuildSupport()
}

// ConcreteImplementorA 是具体的支撑结构A
type ConcreteImplementorA struct{}

func (c *ConcreteImplementorA) BuildSupport() {
	fmt.Println("建造支撑结构A")
}

// ConcreteImplementorB 是具体的支撑结构B
type ConcreteImplementorB struct{}

func (c *ConcreteImplementorB) BuildSupport() {
	fmt.Println("建造支撑结构B")
}

// Abstraction 是抽象部分的接口，定义了功能或特性的方法
type Abstraction interface {
	BuildRoad()
}

// Road 是具体的路（抽象部分），它包含一个实现部分的引用
type Road struct {
	implementor Implementor
}

func NewRoad(implementor Implementor) *Road {
	return &Road{implementor}
}

func (r *Road) BuildRoad() {
	fmt.Println("建造路")
	r.implementor.BuildSupport()
}

func main() {
	// 使用支撑结构A的路
	roadA := NewRoad(&ConcreteImplementorA{})
	roadA.BuildRoad()

	// 使用支撑结构B的路
	roadB := NewRoad(&ConcreteImplementorB{})
	roadB.BuildRoad()
}
