package main

import (
	"fmt"
)

// Shape 是一个图形的接口
type Shape interface {
	Draw() string
}

// Rectangle 是一个矩形结构体，实现了 Shape 接口
type Rectangle struct{}

func (r Rectangle) Draw() string {
	return "Drawing a rectangle"
}

// Circle 是一个圆形结构体，实现了 Shape 接口
type Circle struct{}

func (c Circle) Draw() string {
	return "Drawing a circle"
}

// ShapeFactory 是一个图形工厂接口
type ShapeFactory interface {
	CreateShape() Shape
}

// RectangleFactory 是一个矩形工厂结构体，实现了 ShapeFactory 接口
type RectangleFactory struct{}

func (rf RectangleFactory) CreateShape() Shape {
	return Rectangle{}
}

// CircleFactory 是一个圆形工厂结构体，实现了 ShapeFactory 接口
type CircleFactory struct{}

func (cf CircleFactory) CreateShape() Shape {
	return Circle{}
}

func main() {
	rectangleFactory := RectangleFactory{}
	circleFactory := CircleFactory{}

	rectangle := rectangleFactory.CreateShape()
	circle := circleFactory.CreateShape()

	fmt.Println(rectangle.Draw()) // 输出: Drawing a rectangle
	fmt.Println(circle.Draw())    // 输出: Drawing a circle
}
