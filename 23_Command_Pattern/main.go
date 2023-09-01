package main

import "fmt"

// Command 是命令接口
type Command interface {
	Execute()
}

// Light 是接收者
type Light struct {
	isOn bool
}

func NewLight() *Light {
	return &Light{}
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Println("灯亮了")
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Println("灯灭了")
}

// LightOnCommand 是具体的命令，表示打开灯
type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (c *LightOnCommand) Execute() {
	c.light.TurnOn()
}

// LightOffCommand 是具体的命令，表示关闭灯
type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (c *LightOffCommand) Execute() {
	c.light.TurnOff()
}

// RemoteControl 是调用者
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func main() {
	light := NewLight()
	lightOnCommand := NewLightOnCommand(light)
	lightOffCommand := NewLightOffCommand(light)

	remoteControl := &RemoteControl{}

	remoteControl.SetCommand(lightOnCommand)
	remoteControl.PressButton()

	remoteControl.SetCommand(lightOffCommand)
	remoteControl.PressButton()
}
