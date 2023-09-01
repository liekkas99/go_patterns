package main

import "fmt"

// CoffeeComponent 是咖啡组件接口
type CoffeeComponent interface {
	MakeCoffee() string
}

// SimpleCoffee 是咖啡的基本组件
type SimpleCoffee struct{}

func (c *SimpleCoffee) MakeCoffee() string {
	return "Making a simple coffee"
}

// MilkDecorator 是牛奶装饰类
type MilkDecorator struct {
	component CoffeeComponent
}

func (d *MilkDecorator) MakeCoffee() string {
	return d.component.MakeCoffee() + ", with milk"
}

// SugarDecorator 是糖装饰类
type SugarDecorator struct {
	component CoffeeComponent
}

func (d *SugarDecorator) MakeCoffee() string {
	return d.component.MakeCoffee() + ", with sugar"
}

func main() {
	simpleCoffee := &SimpleCoffee{}
	fmt.Println(simpleCoffee.MakeCoffee())

	milkCoffee := &MilkDecorator{component: simpleCoffee}
	fmt.Println(milkCoffee.MakeCoffee())

	sugarMilkCoffee := &SugarDecorator{component: milkCoffee}
	fmt.Println(sugarMilkCoffee.MakeCoffee())
}
