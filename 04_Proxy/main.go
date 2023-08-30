package main

import "fmt"

// Subject 是主题接口，定义了任务的方法
type Subject interface {
	DoTask()
}

// RealSubject 是真实主题，执行实际任务的对象
type RealSubject struct{}

func (r RealSubject) DoTask() {
	fmt.Println("RealSubject is doing the task.")
}

// Proxy 是代理对象，控制对真实主题的访问
type Proxy struct {
	realSubject RealSubject
}

func (p Proxy) DoTask() {
	fmt.Println("Proxy is doing some preparation before calling the real task.")
	p.realSubject.DoTask()
	fmt.Println("Proxy is doing some cleanup after calling the real task.")
}

func main() {
	proxy := Proxy{}
	proxy.DoTask()
}
