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

	return &btns
}

func (this *BtNodeSelector) Evaluate() bool {
	if len(this.children) == 0 {
		return false
	}

	for i, child := range this.children {
		//fmt.Printf("child %v, time: %v\n", child.GetId(), time.Now().UnixNano())

		//todo:node.Evaluate
		if node.Evaluate(child) {
			//fmt.Printf("child %v, time: %v\n", child.GetTypes(), time.Now().UnixNano())
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

	//fmt.Printf("this.activeChild: %v, childRst: %v\n", this.activeChild.GetTypes(), childRst)

	//运行中
	if childRst == def.Running {
		//fmt.Printf("btni: %v, running\n", this.GetTypes())
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
	//fmt.Printf("reset ------ time: %v\n", time.Now().UnixNano())

	this.SetStatus(def.Ready)
	this.activeIdx = -1
	this.activeChild = nil
	for _, child := range this.children {
		child.Reset()
	}
}
