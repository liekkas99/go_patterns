package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Data string
}

var instance *Singleton
var once sync.Once

// 懒汉式是一种延迟实例化的方式，即在首次使用单例对象时才创建实例
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{Data: "Hello, I am a Singleton instance (Lazy initialization)"}
	})
	return instance
}

func main() {
	singleton1 := GetInstance()
	singleton2 := GetInstance()

	fmt.Println(singleton1 == singleton2) // 输出: true，两个变量引用的是同一个实例
	fmt.Println(singleton1.Data)          // 输出: Hello, I am a Singleton instance (Lazy initialization)
	fmt.Println(singleton2.Data)          // 输出: Hello, I am a Singleton instance (Lazy initialization)
}
