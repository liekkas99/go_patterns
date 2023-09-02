package main

import (
	"fmt"
)

// Flyweight 是享元接口
type Flyweight interface {
	Operation(state string)
}

// ConcreteFlyweight 是具体享元类
type ConcreteFlyweight struct {
	intrinsicState string
}

func NewConcreteFlyweight(intrinsicState string) *ConcreteFlyweight {
	return &ConcreteFlyweight{intrinsicState: intrinsicState}
}

func (f *ConcreteFlyweight) Operation(state string) {
	fmt.Printf("ConcreteFlyweight with intrinsic state %s and extrinsic state %s\n", f.intrinsicState, state)
}

// FlyweightFactory 是享元工厂类
type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{
		flyweights: make(map[string]Flyweight),
	}
}

func (factory *FlyweightFactory) GetFlyweight(key string) Flyweight {
	if flyweight, ok := factory.flyweights[key]; ok {
		return flyweight
	} else {
		flyweight := NewConcreteFlyweight(key)
		factory.flyweights[key] = flyweight
		return flyweight
	}
}

func main() {
	factory := NewFlyweightFactory()

	flyweight1 := factory.GetFlyweight("A")
	flyweight1.Operation("X")

	flyweight2 := factory.GetFlyweight("A")
	flyweight2.Operation("Y")

	flyweight3 := factory.GetFlyweight("B")
	flyweight3.Operation("Z")
}
