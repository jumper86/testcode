package logic_node

import (
	"fmt"
	"testcode/behavior_tree/def"
	"testcode/behavior_tree/node"
)

type BtNodeSequence struct {
	BtLogicNode
	activeIdx   int               //当前执行子节点idx
	activeChild node.BtNodeInterf //当前执行子节点
}

func NewBtNodeSequence(name string, interval int64) *BtNodeSequence {
	var btns BtNodeSequence
	btns.BtNode = node.NewBtNode(name, interval)
	btns.children = make([]node.BtNodeInterf, 0)
	btns.activeIdx = -1
	btns.activeChild = nil
	btns.SetTypes(def.ComposeSequenceNode)
	btns.SetEvaluate(btns.doEvaluate)
	return &btns
}

//Evaluate 只在开始执行该节点时调用一次
func (this *BtNodeSequence) doEvaluate() bool {
	fmt.Printf("sequence.")

	if this.children == nil || len(this.children) == 0 {
		return false
	}

	for _, child := range this.children {
		if !child.Evaluate() {
			return false
		}
	}

	return true
}

func (this *BtNodeSequence) Tick() def.BtnResult {
	if this.activeChild == nil {
		this.activeIdx = 0
		this.activeChild = this.children[0]
	}

	//执行逻辑
	childRst := node.Process(this.activeChild)

	//结果 运行中
	if childRst == def.Running {
		return def.Running
	}

	//结果 成功
	//note:
	//      在顺序节点中，还有另一种方式，就是在子节点执行成功的时候，继续执行下一个子节点，直到某个子节点返回运行中或者执行失败。
	//      https://www.behaviac.com/concepts/
	//      但是这种方式可能存在问题，可能导致这个顺序节点所花费的时间太长。
	//      而这里当执行一个子节点成功了就立即返回运行中，能够防止该顺序节点花费过长时间

	if childRst == def.Successed {
		if this.activeIdx == len(this.children)-1 {
			return def.Successed
		} else {
			this.activeIdx++
			this.activeChild = this.children[this.activeIdx]
			return def.Running
		}
	}

	//结果 失败
	if childRst == def.Failed {
		return def.Failed
	}

	return def.Failed
}

//准入失败时，执行成功，执行失败时调用
func (this *BtNodeSequence) Reset() {
	this.SetStatus(def.Ready)
	this.activeIdx = -1
	this.activeChild = nil
	for _, child := range this.children {
		child.Reset()
	}
}
