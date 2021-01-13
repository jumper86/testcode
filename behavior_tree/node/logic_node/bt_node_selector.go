package logic_node

import (
	"testcode/behavior_tree/def"
	"testcode/behavior_tree/node"
)

type BtNodeSelector struct {
	BtLogicNode
	activeIdx   int
	activeChild node.BtNodeInterf //当前执行子节点
}

func NewBtNodeSelector(name string, interval int64) *BtNodeSelector {
	var btns BtNodeSelector
	btns.BtNode = node.NewBtNode(name, interval)
	btns.children = make([]node.BtNodeInterf, 0)
	btns.activeIdx = -1
	btns.activeChild = nil
	btns.SetTypes(def.ComposeSelectorNode)
	btns.SetEvaluate(btns.doEvaluate)

	return &btns
}

func (this *BtNodeSelector) doEvaluate() bool {
	if len(this.children) == 0 {
		return false
	}

	for i, child := range this.children {
		if child.Evaluate() {
			this.activeIdx = i
			this.activeChild = child
			return true
		}
	}

	return false
}

func (this *BtNodeSelector) Tick() def.BtnResult {

	if this.activeChild == nil {
		return def.Failed
	}

	childRst := this.activeChild.Process()

	//运行中
	if childRst == def.Running {
		return def.Running
	}

	//成功
	if childRst == def.Successed {
		return def.Successed
	}

	//失败
	if childRst == def.Failed {
		//寻找下一个可以执行的子节点
		found := false
		for i := this.activeIdx; i < len(this.children); i++ {
			tmp := this.children[i]
			if tmp.Evaluate() {
				found = true
				this.activeIdx = i
				this.activeChild = tmp
				break
			}
		}

		if !found {
			return def.Failed
		} else {
			return def.Running
		}
	}

	return def.Failed
}

func (this *BtNodeSelector) Reset() {
	this.SetStatus(def.Ready)
	this.activeIdx = -1
	this.activeChild = nil
	for _, child := range this.children {
		child.Reset()
	}
}

func (this *BtNodeSelector) Process() def.BtnResult {
	if !this.Evaluate() {
		return def.Failed
	}
	if this.GetStatus() != def.Run {
		this.SetStatus(def.Run)
	}

	tmpRst := this.Tick()
	if tmpRst != def.Running {
		this.Reset()
	}
	return tmpRst
}
