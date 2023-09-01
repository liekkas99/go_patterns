package main

import "fmt"

// Component 是组合模式中的基本组件接口
type Component interface {
	Operation() string
}

// Leaf 是叶子节点，实现了 Component 接口
type Leaf struct {
	name string
}

func (l *Leaf) Operation() string {
	return fmt.Sprintf("Leaf %s", l.name)
}

// Composite 是组合对象，包含 Leaf 和其他 Composite 对象
type Composite struct {
	name       string
	components []Component
}

func (c *Composite) Operation() string {
	result := fmt.Sprintf("Composite %s [", c.name)
	for _, component := range c.components {
		result += component.Operation() + " "
	}
	result += "]"
	return result
}

func main() {
	leaf1 := &Leaf{name: "Leaf 1"}
	leaf2 := &Leaf{name: "Leaf 2"}
	composite1 := &Composite{name: "Composite 1", components: []Component{leaf1, leaf2}}

	leaf3 := &Leaf{name: "Leaf 3"}
	leaf4 := &Leaf{name: "Leaf 4"}
	composite2 := &Composite{name: "Composite 2", components: []Component{leaf3, leaf4, composite1}}

	fmt.Println(composite2.Operation())
}
