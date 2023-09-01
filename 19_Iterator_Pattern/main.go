package main

import (
	"fmt"
)

// Iterator 定义迭代器接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// Aggregate 定义集合接口
type Aggregate interface {
	CreateIterator() Iterator
}

// ConcreteIterator 实现迭代器接口
type ConcreteIterator struct {
	aggregate *ConcreteAggregate
	index     int
}

func NewConcreteIterator(aggregate *ConcreteAggregate) *ConcreteIterator {
	return &ConcreteIterator{
		aggregate: aggregate,
		index:     0,
	}
}

func (it *ConcreteIterator) HasNext() bool {
	return it.index < len(it.aggregate.items)
}

func (it *ConcreteIterator) Next() interface{} {
	item := it.aggregate.items[it.index]
	it.index++
	return item
}

// ConcreteAggregate 实现集合接口
type ConcreteAggregate struct {
	items []interface{}
}

func NewConcreteAggregate() *ConcreteAggregate {
	return &ConcreteAggregate{
		items: make([]interface{}, 0),
	}
}

func (a *ConcreteAggregate) CreateIterator() Iterator {
	return NewConcreteIterator(a)
}

func (a *ConcreteAggregate) AddItem(item interface{}) {
	a.items = append(a.items, item)
}

func main() {
	aggregate := NewConcreteAggregate()
	aggregate.AddItem("Item 1")
	aggregate.AddItem("Item 2")
	aggregate.AddItem("Item 3")

	iterator := aggregate.CreateIterator()

	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Println(item)
	}
}
