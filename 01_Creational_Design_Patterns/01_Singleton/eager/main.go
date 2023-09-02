package main

import (
	"fmt"
)

type Singleton struct {
	Data string
}

// 饿汉式是一种在程序启动时就立即创建实例的方式
var instance *Singleton = &Singleton{Data: "Hello, I am a Singleton instance (Eager initialization)"}

func GetInstance() *Singleton {
	return instance
}

func main() {
	singleton1 := GetInstance()
	singleton2 := GetInstance()

	fmt.Println(singleton1 == singleton2) // 输出: true，两个变量引用的是同一个实例
	fmt.Println(singleton1.Data)          // 输出: Hello, I am a Singleton instance (Eager initialization)
	fmt.Println(singleton2.Data)          // 输出: Hello, I am a Singleton instance (Eager initialization)
}
