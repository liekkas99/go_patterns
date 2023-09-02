package main

import "fmt"

// Target 是目标接口，定义了客户端期望使用的方法
type Target interface {
	Request()
}

// Adaptee 是被适配的类，拥有不兼容的方法
type Adaptee struct{}

func (a Adaptee) SpecificRequest() {
	fmt.Println("SpecificRequest from Adaptee")
}

// Adapter 是适配器，将 Adaptee 的方法适配成 Target 的方法
type Adapter struct {
	adaptee Adaptee
}

func (a Adapter) Request() {
	fmt.Println("Request from Adapter")
	a.adaptee.SpecificRequest()
}

func main() {
	adaptee := Adaptee{}
	adapter := Adapter{adaptee}

	adapter.Request()
}
