package main

import "fmt"

// Skill 是技能接口，定义了不同技能的执行方法
type Skill interface {
	Execute()
}

// Warrior 是战士角色，实现了技能接口
type Warrior struct{}

func (w *Warrior) Execute() {
	fmt.Println("战士使用近身攻击")
}

// Mage 是法师角色，实现了技能接口
type Mage struct{}

func (m *Mage) Execute() {
	fmt.Println("法师施放魔法")
}

// Archer 是射手角色，实现了技能接口
type Archer struct{}

func (a *Archer) Execute() {
	fmt.Println("射手进行远程攻击")
}

// Player 是玩家，具有切换角色和执行技能的能力
type Player struct {
	currentSkill Skill
}

func (p *Player) SetSkill(skill Skill) {
	p.currentSkill = skill
}

func (p *Player) UseSkill() {
	if p.currentSkill == nil {
		fmt.Println("请选择一个角色和技能")
		return
	}

	p.currentSkill.Execute()
}

func main() {
	player := &Player{}

	// 选择战士角色并使用技能
	player.SetSkill(&Warrior{})
	player.UseSkill()

	// 选择法师角色并使用技能
	player.SetSkill(&Mage{})
	player.UseSkill()

	// 选择射手角色并使用技能
	player.SetSkill(&Archer{})
	player.UseSkill()
}
