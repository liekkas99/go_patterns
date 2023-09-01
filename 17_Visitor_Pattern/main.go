package main

import "fmt"

// Visitor 是访问者接口
type Visitor interface {
	VisitAnimal(animal Animal)
}

// Animal 是被访问者接口
type Animal interface {
	Accept(visitor Visitor)
}

// Lion 是具体的被访问者
type Lion struct{}

func (l *Lion) Accept(visitor Visitor) {
	visitor.VisitAnimal(l)
}

// Tiger 是具体的被访问者
type Tiger struct{}

func (t *Tiger) Accept(visitor Visitor) {
	visitor.VisitAnimal(t)
}

// AnimalOperation 是具体的访问者
type AnimalOperation struct{}

func (ao *AnimalOperation) VisitAnimal(animal Animal) {
	switch animal.(type) {
	case *Lion:
		fmt.Println("Feeding lion")
	case *Tiger:
		fmt.Println("Feeding tiger")
	}
}

func main() {
	lion := &Lion{}
	tiger := &Tiger{}

	operation := &AnimalOperation{}

	lion.Accept(operation)  // 输出: Feeding lion
	tiger.Accept(operation) // 输出: Feeding tiger
}
