package main

import (
	"fmt"
	"time"
)

// Observer 是观察者接口
type Observer interface {
	Update(news string)
}

// Subject 是被观察者接口
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify(news string)
}

// NewsPublisher 是新闻发布者，实现了 Subject 接口
type NewsPublisher struct {
	observers []Observer
}

func NewNewsPublisher() *NewsPublisher {
	return &NewsPublisher{}
}

func (np *NewsPublisher) Register(observer Observer) {
	np.observers = append(np.observers, observer)
}

func (np *NewsPublisher) Unregister(observer Observer) {
	for i, obs := range np.observers {
		if obs == observer {
			np.observers = append(np.observers[:i], np.observers[i+1:]...)
			break
		}
	}
}

func (np *NewsPublisher) Notify(news string) {
	for _, observer := range np.observers {
		observer.Update(news)
	}
}

// NewsSubscriber 是新闻订阅者，实现了 Observer 接口
type NewsSubscriber struct {
	name string
}

func NewNewsSubscriber(name string) *NewsSubscriber {
	return &NewsSubscriber{name}
}

func (ns *NewsSubscriber) Update(news string) {
	fmt.Printf("%s 收到新闻：%s\n", ns.name, news)
}

func main() {
	newsPublisher := NewNewsPublisher()

	subscriber1 := NewNewsSubscriber("订阅者1")
	subscriber2 := NewNewsSubscriber("订阅者2")

	newsPublisher.Register(subscriber1)
	newsPublisher.Register(subscriber2)

	// 模拟新闻发布
	newsPublisher.Notify("重要新闻：今天是个好天气！")

	// 移除一个订阅者
	newsPublisher.Unregister(subscriber1)

	// 模拟新闻发布
	newsPublisher.Notify("紧急新闻：交通事故发生！")

	time.Sleep(1 * time.Second)
}
