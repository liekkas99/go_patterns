package main

import "fmt"

// Request 表示请求的结构体
type Request struct {
	name     string
	numDays  int
	reason   string
	approved bool
}

// Approver 是审批者接口
type Approver interface {
	SetNext(Approver)
	ProcessRequest(Request)
}

// Supervisor 是主管结构体
type Supervisor struct {
	next Approver
	name string
}

func (s *Supervisor) SetNext(next Approver) {
	s.next = next
}

func (s *Supervisor) ProcessRequest(request Request) {
	if request.numDays <= 2 {
		fmt.Printf("主管 %s 批准了请假申请：%s\n", s.name, request.reason)
		request.approved = true
	} else if s.next != nil {
		s.next.ProcessRequest(request)
	}
}

// Manager 是经理结构体
type Manager struct {
	next Approver
	name string
}

func (m *Manager) SetNext(next Approver) {
	m.next = next
}

func (m *Manager) ProcessRequest(request Request) {
	if request.numDays <= 5 {
		fmt.Printf("经理 %s 批准了请假申请：%s\n", m.name, request.reason)
		request.approved = true
	} else if m.next != nil {
		m.next.ProcessRequest(request)
	}
}

// Director 是总监结构体
type Director struct {
	next Approver
	name string
}

func (d *Director) SetNext(next Approver) {
	d.next = next
}

func (d *Director) ProcessRequest(request Request) {
	if request.numDays <= 7 {
		fmt.Printf("总监 %s 批准了请假申请：%s\n", d.name, request.reason)
		request.approved = true
	} else {
		fmt.Printf("申请被拒绝：%s\n", request.reason)
	}
}

func main() {
	director := &Director{name: "张总监"}
	manager := &Manager{name: "李经理"}
	supervisor := &Supervisor{name: "王主管"}

	director.SetNext(manager)
	manager.SetNext(supervisor)

	request := Request{
		name:     "小明",
		numDays:  4,
		reason:   "家里有事情",
		approved: false,
	}

	director.ProcessRequest(request)
	if request.approved {
		fmt.Println("申请已批准")
	} else {
		fmt.Println("申请被拒绝")
	}
}

/*
在这个示例中，我们定义了 `Request` 结构体表示请求，
以及 `Approver` 接口表示审批者。然后，我们实现了
具体的审批者：`Supervisor`、`Manager` 和 `Director`。
每个审批者都有一个 `SetNext` 方法来设置下一个处理者，
以及一个 `ProcessRequest` 方法来处理请求。如果当前审批者无法处理请求
它会将请求传递给下一个审批者。最终，我们创建了一个责任链，
从最高层的 `Director` 开始，逐级传递请求，直到找到可以处理请求的审批者或整个链条结束。
*/
