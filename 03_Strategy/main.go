package main

import "fmt"

// Strategy 是策略接口，定义了执行算法的方法
type Strategy interface {
	Execute()
}

// ConcreteStrategyA 是具体策略A
type ConcreteStrategyA struct{}

func (s ConcreteStrategyA) Execute() {
	fmt.Println("Executing ConcreteStrategyA")
}

// ConcreteStrategyB 是具体策略B
type ConcreteStrategyB struct{}

func (s ConcreteStrategyB) Execute() {
	fmt.Println("Executing ConcreteStrategyB")
}

// Context 是上下文类，持有一个策略对象
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy() {
	if c.strategy != nil {
		c.strategy.Execute()
	}
}

func main() {
	context := Context{}

	strategyA := ConcreteStrategyA{}
	context.SetStrategy(strategyA)
	context.ExecuteStrategy() // 输出: Executing ConcreteStrategyA

	strategyB := ConcreteStrategyB{}
	context.SetStrategy(strategyB)
	context.ExecuteStrategy() // 输出: Executing ConcreteStrategyB
}
