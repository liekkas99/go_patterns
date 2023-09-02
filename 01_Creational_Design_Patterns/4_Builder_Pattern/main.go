package main

import "fmt"

// Product 表示最终构建出的产品
type Product struct {
	mainCourse string
	sideDish   string
	drink      string
}

// Builder 是抽象的建造者接口
type Builder interface {
	BuildMainCourse()
	BuildSideDish()
	BuildDrink()
	GetProduct() Product
}

// Director 是指挥者，负责指导建造的过程
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() Product {
	d.builder.BuildMainCourse()
	d.builder.BuildSideDish()
	d.builder.BuildDrink()
	return d.builder.GetProduct()
}

// FastFoodBuilder 是具体的建造者
type FastFoodBuilder struct {
	product Product
}

func NewFastFoodBuilder() *FastFoodBuilder {
	return &FastFoodBuilder{product: Product{}}
}

func (b *FastFoodBuilder) BuildMainCourse() {
	b.product.mainCourse = "汉堡"
}

func (b *FastFoodBuilder) BuildSideDish() {
	b.product.sideDish = "薯条"
}

func (b *FastFoodBuilder) BuildDrink() {
	b.product.drink = "可乐"
}

func (b *FastFoodBuilder) GetProduct() Product {
	return b.product
}

func main() {
	builder := NewFastFoodBuilder()
	director := NewDirector(builder) // 构建
	meal := director.Construct()

	fmt.Println("主菜:", meal.mainCourse)
	fmt.Println("配菜:", meal.sideDish)
	fmt.Println("饮料:", meal.drink)
}

/*
在上述示例中，FastFoodBuilder 负责构建产品的细节，而 Director 则根据需要来指导构建过程。
产品的表示形式由 Product 结构体来定义，它包含了主菜、配菜和饮料等属性。
*/
