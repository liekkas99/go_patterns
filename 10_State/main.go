/*
在这个示例中，我们首先定义了状态接口 State，
其中包含一个 Handle() 方法，用于在不同状态下处理对象的行为。
然后，我们实现了两个具体状态类 ConcreteStateA 和 ConcreteStateB，
分别实现了不同状态下的行为。

接下来，我们定义了上下文类 Context，它持有一个状态对象，
并通过 SetState() 方法来设置不同的状态。Request() 方法用于触发当前状态的行为处理。
在 main 函数中，我们创建了上下文对象，并设置初始状态为 ConcreteStateA。
通过多次调用 Request() 方法，我们可以看到对象的状态在不同情况下改变，
从而表现出不同的行为。这就是状态模式的核心思想在 Go 中的实现方式。
*/package main

import "fmt"

// State 是状态接口，定义了不同状态下的方法
type State interface {
	Handle(context *Context)
}

// ConcreteStateA 是具体状态A
type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle(context *Context) {
	fmt.Println("Handling in State A")
}

// ConcreteStateB 是具体状态B
type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle(context *Context) {
	fmt.Println("Handling in State B")
}

// Context 是上下文类，持有一个状态对象
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle(c)
}

func main() {
	context := Context{}
	initialState := &ConcreteStateA{}
	context.SetState(initialState)

	context.Request() // 输出: Handling in State A

	context.SetState(&ConcreteStateB{})
	context.Request() // 输出: Handling in State B

	context.SetState(&ConcreteStateA{})
	context.Request() // 输出: Handling in State A
}
