package behavior_tree

import "fmt"

type BtNodeSequence struct {
	BtNode
	children    []BtNodeInterf //所有子节点
	activeIdx   int            //当前执行子节点idx
	activeChild BtNodeInterf   //当前执行子节点
}

func NewBtNodeSequence(name string, interval int64) BtNodeSequence {
	var btns BtNodeSequence
	btns.BtNode = NewBtNode(name, interval)
	btns.children = make([]BtNodeInterf, 0)
	btns.activeIdx = -1
	btns.activeChild = nil
	btns.types = ComposeSequenceNode
	return btns
}

func (this *BtNodeSequence) Evaluate() bool {
	//note: 保证只在第一次执行组合节点的时候，进行一次准入检查，即调用 Evaluate
	if this.status != Ready {
		return true
	}
	if this.activated && this.CheckTimer() && this.DoEvaluate() {
		this.status = Running
		return true
	}
	return false
}

//Evaluate 只在开始执行该节点时调用一次
func (this *BtNodeSequence) DoEvaluate() bool {
	fmt.Printf("sequence.")
	if len(this.children) == 0 {
		return false
	}

	for _, child := range this.children {
		if !child.Evaluate() {
			return false
		}
	}

	return true
}

func (this *BtNodeSequence) Tick() BtnResult {
	if this.activeChild == nil {
		this.activeIdx = 0
		this.activeChild = this.children[0]
	}

	//执行逻辑
	childRst := this.activeChild.Tick()

	//结果 运行中
	if childRst == Running {
		return Running
	}

	//结果 成功
	if childRst == Successed {
		if this.activeIdx == len(this.children)-1 {
			this.Reset()
			return Successed
		} else {
			this.activeIdx++
			this.activeChild = this.children[this.activeIdx]
			return Running
		}
	}

	//结果 失败
	if childRst == Failed {
		this.Reset()
		return Failed
	}

	return Failed
}

//准入失败时，执行成功，执行失败时调用
func (this *BtNodeSequence) Reset() {
	this.status = Ready
	this.activeIdx = -1
	this.activeChild = nil
	for _, child := range this.children {
		child.Reset()
	}
}

func (this *BtNodeSequence) AddChild(bn BtNodeInterf) {
	if this.children == nil {
		this.children = make([]BtNodeInterf, 0)
	}
	if bn != nil {
		this.children = append(this.children, bn)
	}

	this.Reset()
	return
}

func (this *BtNodeSequence) RemoveChild(bn BtNodeInterf) {
	objId := bn.GetId()
	objIdx := -1
	for idx, child := range this.children {
		if child.GetId() == objId {
			objIdx = idx
			break
		}
	}
	if objIdx != -1 {
		this.children = append(this.children[:objIdx], this.children[objIdx+1:]...)
	}

	this.Reset()
	return
}

//
//func (this *BtNodeSequence) Process() BtnResult {
//	if this.Evaluate() {
//		return this.Tick()
//	}
//	return Failed
//}
