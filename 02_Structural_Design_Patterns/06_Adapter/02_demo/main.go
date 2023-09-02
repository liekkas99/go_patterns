/*
原本要用中式插座，但是插座是英式的。
然後在中式和英式之間加了個適配器，
將英式插座轉爲了中式插座
*/
package main

import "fmt"

// 中式插座接口
type CCnOutlet interface {
	Cnplug()
}

// 英式插座
type CEnOutlet struct{}

func (e CEnOutlet) EnPlug() {
	fmt.Println("use en plug")
}

// 中式插座适配器
type CCnOutletAdapter struct {
	enOutlet *CEnOutlet
}

func NewCCnOutletAdapter(enOutlet *CEnOutlet) CCnOutlet {
	return &CCnOutletAdapter{enOutlet: enOutlet}
}

func (a *CCnOutletAdapter) Cnplug() {
	fmt.Println("adapter transfer")
	a.enOutlet.EnPlug()
}

func main() {
	enOutlet := &CEnOutlet{}
	cnOutlet := NewCCnOutletAdapter(enOutlet)
	cnOutlet.Cnplug()
}
