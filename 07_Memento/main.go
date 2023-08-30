package main

import "fmt"

// Memento 是备忘录接口，定义了保存和恢复状态的方法
type Memento interface {
	GetState() string
}

// Originator 是原发器，拥有需要保存的状态
type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) CreateMemento() Memento {
	return &MementoImpl{state: o.state}
}

func (o *Originator) RestoreMemento(m Memento) {
	o.state = m.GetState()
}

// MementoImpl 是备忘录实现
type MementoImpl struct {
	state string
}

func (m *MementoImpl) GetState() string {
	return m.state
}

func main() {
	originator := &Originator{}
	originator.SetState("State 1")
	fmt.Println("Current state:", originator.GetState())

	// 保存状态
	memento := originator.CreateMemento()

	originator.SetState("State 2")
	fmt.Println("Current state:", originator.GetState())

	// 恢复状态
	originator.RestoreMemento(memento)
	fmt.Println("Restored state:", originator.GetState())
}
