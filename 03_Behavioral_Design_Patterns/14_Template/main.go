package main

import "fmt"

// Template 是模板接口，定义了算法的基本步骤
type Template interface {
	Prepare()
	Mix()
	Bake()
	Decorate()
}

// Cake 是具体模板，实现了模板接口
type Cake struct{}

func (c Cake) Prepare() {
	fmt.Println("Preparing the cake ingredients")
}

func (c Cake) Mix() {
	fmt.Println("Mixing the cake batter")
}

func (c Cake) Bake() {
	fmt.Println("Baking the cake")
}

func (c Cake) Decorate() {
	fmt.Println("Decorating the cake")
}

// Cupcake 是另一个具体模板
type Cupcake struct{}

func (c Cupcake) Prepare() {
	fmt.Println("Preparing the cupcake ingredients")
}

func (c Cupcake) Mix() {
	fmt.Println("Mixing the cupcake batter")
}

func (c Cupcake) Bake() {
	fmt.Println("Baking the cupcake")
}

func (c Cupcake) Decorate() {
	fmt.Println("Decorating the cupcake")
}

func main() {
	cake := Cake{}
	cupcake := Cupcake{}

	fmt.Println("Making a cake:")
	cake.Prepare()
	cake.Mix()
	cake.Bake()
	cake.Decorate()

	fmt.Println("\nMaking a cupcake:")
	cupcake.Prepare()
	cupcake.Mix()
	cupcake.Bake()
	cupcake.Decorate()
}
