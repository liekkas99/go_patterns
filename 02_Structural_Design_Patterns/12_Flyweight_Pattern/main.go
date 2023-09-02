package main

import (
	"fmt"
)

// ToyFactory 是玩具工厂，负责创建玩具
type ToyFactory struct {
	toys map[string]Toy
}

func NewToyFactory() *ToyFactory {
	return &ToyFactory{
		toys: make(map[string]Toy),
	}
}

// Toy 是玩具接口，定义了展示玩具的方法
type Toy interface {
	Show()
}

// Car 是具体的玩具类，表示小汽车玩具
type Car struct {
	color string
}

func (c *Car) Show() {
	fmt.Printf("这是一辆%s小汽车\n", c.color)
}

// Robot 是具体的玩具类，表示机器人玩具
type Robot struct {
	model string
}

func (r *Robot) Show() {
	fmt.Printf("这是一个%s型号的机器人\n", r.model)
}

// GetToy 根据玩具类型获取玩具，如果已有相同类型的玩具，则共享
func (f *ToyFactory) GetToy(toyType string) Toy {
	if toy, ok := f.toys[toyType]; ok {
		fmt.Println("共享相同类型的玩具")
		return toy
	}

	var newToy Toy
	switch toyType {
	case "小汽车":
		newToy = &Car{color: "红色"}
	case "机器人":
		newToy = &Robot{model: "机器人模型-1"}
	default:
		fmt.Println("不支持的玩具类型")
		return nil
	}

	f.toys[toyType] = newToy
	return newToy
}

func main() {
	factory := NewToyFactory()

	// 创建小汽车玩具，并展示
	car1 := factory.GetToy("小汽车")
	car1.Show()

	// 再次创建小汽车玩具，此时应共享相同类型的玩具
	car2 := factory.GetToy("小汽车")
	car2.Show()

	// 创建机器人玩具
	robot := factory.GetToy("机器人")
	robot.Show()
}
