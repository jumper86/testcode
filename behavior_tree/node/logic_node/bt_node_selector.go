package logic_node

import (
	"testcode/behavior_tree/def"
	"testcode/behavior_tree/node"
)

//note:
//	检查方式：子节点中有一个满足，则检查通过，并且会选定当前子节点
//	执行方式：从检查过程中选定的子节点开始执行，若是失败就选择下一个可执行的子节点，直到某个子节点成功。每次Tick执行一次当前子节点的Process
//	返回方式：
//		子节点返回running，则返回running
//		子节点返回successed，则返回successed
//		非最后子节点返回failed，则返回running；最后子节点返回failed，则返回failed

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

	return &btns
}

func (this *BtNodeSelector) Evaluate() bool {
	if len(this.children) == 0 {
		return false
	}

	for i, child := range this.children {

		if node.Evaluate(child) {
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

	childRst := node.Process(this.activeChild)

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
			if node.Evaluate(tmp) {
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
