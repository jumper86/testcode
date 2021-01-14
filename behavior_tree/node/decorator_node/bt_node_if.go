package decorator_node

import (
	"testcode/behavior_tree/def"
	"testcode/behavior_tree/node"
)

//note:
//	运行子节点之前需要添加函数进行过滤
type BtNodeIf struct {
	BtDecoratorNode
	cond  func() bool       //条件
	child node.BtNodeInterf //所有子节点
}

func NewBtNodeIf(name string, interval int64, cond func() bool) *BtNodeIf {
	var btns BtNodeIf
	btns.BtNode = node.NewBtNode(name, interval)
	btns.SetTypes(def.DecoratorIfNode)
	btns.cond = cond
	return &btns
}

func (this *BtNodeIf) Evaluate() bool {
	if !this.cond() {
		return false
	}
	return node.Evaluate(this.child)
}

func (this *BtNodeIf) Tick() def.BtnResult {

	if !this.cond() {
		return def.Failed
	}

	childRst := node.Process(this.child)
	return childRst
}
