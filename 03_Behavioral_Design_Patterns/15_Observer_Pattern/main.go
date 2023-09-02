package main

import "fmt"

// Observer 定义观察者接口
type Observer interface {
	Update(message string)
}

// Subject 定义被观察的主题接口
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(message string)
}

// ConcreteObserver 实现了 Observer 接口
type ConcreteObserver struct {
	name string
}

func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("%s 收到通知：%s\n", o.name, message)
}

// ConcreteSubject 实现了 Subject 接口
type ConcreteSubject struct {
	observers []Observer
}

func (s *ConcreteSubject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyObservers(message string) {
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

func main() {
	subject := &ConcreteSubject{}

	observer1 := &ConcreteObserver{name: "观察者1"}
	observer2 := &ConcreteObserver{name: "观察者2"}

	subject.RegisterObserver(observer1)
	subject.RegisterObserver(observer2)

	subject.NotifyObservers("新消息：Hello, 观察者们！")

	subject.RemoveObserver(observer1)

	subject.NotifyObservers("又有新消息：大家注意了！")
}
