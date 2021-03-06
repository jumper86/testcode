package logic_node

import (
	"testcode/behavior_tree/def"
	"testcode/behavior_tree/node"
)

//note:
//	检查方式：子节点中有一个不满足，则检查不通过
//	执行方式：每次Tick依次执行当前子节点的Process
//			该节点用于处理如下逻辑组合：
//			a b 两个节点，a 节点执行失败则执行 b 节点，b 节点执行失败则整个组合节点失败，
//			b 节点执行成功，则再次回去执行 a 节点，如此往复。
//			比如 a 节点表示追击，b 节点表示攻击
//	返回方式：
//		最后一个子节点返回failed，则返回failed

type BtNodeAnchorLoop struct {
	BtLogicNode
	activeIdx   int
	activeChild node.BtNodeInterf //当前执行子节点
	anchorIdx   int               //锚点
}

func NewBtNodeAnchorLoop(name string, interval int64, anchorIdx int) *BtNodeAnchorLoop {
	var btns BtNodeAnchorLoop
	btns.BtNode = node.NewBtNode(name, interval)
	btns.children = make([]node.BtNodeInterf, 0)
	btns.activeIdx = -1
	btns.activeChild = nil
	btns.anchorIdx = anchorIdx
	btns.SetTypes(def.ComposeAnchorLoopNode)

	return &btns
}

func (this *BtNodeAnchorLoop) Evaluate() bool {
	l := len(this.children)
	if l == 0 || this.anchorIdx >= l {
		return false
	}

	for _, child := range this.children {
		if !node.Evaluate(child) {
			return false
		}
	}

	this.activeIdx = 0
	this.activeChild = this.children[this.activeIdx]
	return true
}

func (this *BtNodeAnchorLoop) Tick() def.BtnResult {

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
		//note: 这个reset需要保留
		this.Reset()
		this.activeIdx = this.anchorIdx
		this.activeChild = this.children[this.activeIdx]
		return def.Running
	}

	//失败
	if childRst == def.Failed {
		if this.activeIdx == len(this.children)-1 {
			return def.Failed
		} else {
			this.activeIdx++
			this.activeChild = this.children[this.activeIdx]
			return def.Running
		}
	}

	return def.Failed
}

func (this *BtNodeAnchorLoop) Reset() {
	this.SetStatus(def.Ready)
	this.activeIdx = -1
	this.activeChild = nil
	for _, child := range this.children {
		child.Reset()
	}
}
