package main

import "fmt"

// Implementor 是实现者接口
type Implementor interface {
	Operate()
}

// ConcreteImplementorA 是具体实现者A
type ConcreteImplementorA struct{}

func (a *ConcreteImplementorA) Operate() {
	fmt.Println("Concrete Implementor A operation")
}

// ConcreteImplementorB 是具体实现者B
type ConcreteImplementorB struct{}

func (b *ConcreteImplementorB) Operate() {
	fmt.Println("Concrete Implementor B operation")
}

// Abstraction 是抽象类
type Abstraction struct {
	implementor Implementor
}

func (ab *Abstraction) SetImplementor(impl Implementor) {
	ab.implementor = impl
}

func (ab *Abstraction) DoOperation() {
	ab.implementor.Operate()
}

func main() {
	implementorA := &ConcreteImplementorA{}
	implementorB := &ConcreteImplementorB{}

	abstractionA := &Abstraction{}
	abstractionA.SetImplementor(implementorA)
	abstractionA.DoOperation() // 输出: Concrete Implementor A operation

	abstractionB := &Abstraction{}
	abstractionB.SetImplementor(implementorB)
	abstractionB.DoOperation() // 输出: Concrete Implementor B operation
}
