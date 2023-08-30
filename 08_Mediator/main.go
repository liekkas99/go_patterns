package main

import "fmt"

// Mediator 是中介者接口，定义了协调对象之间交互的方法
type Mediator interface {
	SendMessage(user User, message string)
}

// ConcreteMediator 是具体中介者，实现了中介者接口
type ConcreteMediator struct{}

func (m ConcreteMediator) SendMessage(user User, message string) {
	fmt.Printf("%s sends message: %s\n", user.GetName(), message)
}

// User 是参与交互的对象
type User struct {
	name     string
	mediator Mediator
}

func NewUser(name string, mediator Mediator) User {
	return User{name: name, mediator: mediator}
}

func (u User) GetName() string {
	return u.name
}

func (u User) SendMessage(message string) {
	u.mediator.SendMessage(u, message)
}

func main() {
	mediator := ConcreteMediator{}

	user1 := NewUser("User 1", mediator)
	user2 := NewUser("User 2", mediator)
	user3 := NewUser("User 3", mediator)

	user1.SendMessage("Hello, everyone!")
	user2.SendMessage("Hey there!")
	user3.SendMessage("Greetings!")

	// Output:
	// User 1 sends message: Hello, everyone!
	// User 2 sends message: Hey there!
	// User 3 sends message: Greetings!
}
