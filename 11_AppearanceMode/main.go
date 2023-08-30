package main

import "fmt"

// CPU 子系统
type CPU struct{}

func (c *CPU) Start() {
	fmt.Println("CPU is starting")
}

func (c *CPU) Shutdown() {
	fmt.Println("CPU is shutting down")
}

// Memory 子系统
type Memory struct{}

func (m *Memory) Load() {
	fmt.Println("Memory is loading data")
}

func (m *Memory) Unload() {
	fmt.Println("Memory is unloading data")
}

// Computer 外观类，封装了多个子系统的操作
type Computer struct {
	cpu    *CPU
	memory *Memory
}

func NewComputer() *Computer {
	return &Computer{
		cpu:    &CPU{},
		memory: &Memory{},
	}
}

func (c *Computer) Start() {
	c.cpu.Start()
	c.memory.Load()
	fmt.Println("Computer is starting")
}

func (c *Computer) Shutdown() {
	c.memory.Unload()
	c.cpu.Shutdown()
	fmt.Println("Computer is shutting down")
}

func main() {
	computer := NewComputer()

	// 用户只需要调用外观类的方法即可
	computer.Start()
	computer.Shutdown()
}
