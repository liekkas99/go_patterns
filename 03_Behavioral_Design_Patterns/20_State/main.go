package main

import "fmt"

// State 是状态接口，定义了自动售货机状态下的行为
type State interface {
	InsertCoin()
	PushButton()
}

// VendingMachine 是自动售货机，维护当前的状态
type VendingMachine struct {
	state State
}

func NewVendingMachine() *VendingMachine {
	return &VendingMachine{state: &NoStockState{}}
}

func (vm *VendingMachine) SetState(state State) {
	vm.state = state
}

func (vm *VendingMachine) InsertCoin() {
	vm.state.InsertCoin()
}

func (vm *VendingMachine) PushButton() {
	vm.state.PushButton()
}

// NoStockState 是无货状态，实现了 State 接口
type NoStockState struct{}

func (nss *NoStockState) InsertCoin() {
	fmt.Println("无法投币，商品已售罄")
}

func (nss *NoStockState) PushButton() {
	fmt.Println("无法购买，商品已售罄")
}

// InStockState 是有货状态，实现了 State 接口
type InStockState struct{}

func (iss *InStockState) InsertCoin() {
	fmt.Println("投币成功，等待选择商品")
}

func (iss *InStockState) PushButton() {
	fmt.Println("商品已出售，感谢购买")
}

func main() {
	vendingMachine := NewVendingMachine()

	vendingMachine.InsertCoin()
	vendingMachine.PushButton()

	vendingMachine.SetState(&InStockState{})
	vendingMachine.InsertCoin()
	vendingMachine.PushButton()
}
