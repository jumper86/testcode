package decorator_node

import (
	"testcode/behavior_tree/def"
	"testcode/behavior_tree/node"
)

type BtNodeIf struct {
	BtDecoratorNode
	cond  func() bool       //条件
	child node.BtNodeInterf //所有子节点
}

func NewBtNodeIf(name string, interval int64, cond func() bool) *BtNodeIf {
	var btns BtNodeIf
	btns.BtNode = node.NewBtNode(name, interval)
	btns.SetTypes(def.DecoratorIfNode)
	btns.SetEvaluate(btns.doEvaluate)
	btns.cond = cond
	return &btns
}

func (this *BtNodeIf) doEvaluate() bool {
	if !this.cond() {
		return false
	}
	return this.child.Evaluate()
}

func (this *BtNodeIf) Tick() def.BtnResult {

	if !this.cond() {
		return def.Failed
	}

	childRst := node.Process(this.child)
	//结果 运行中
	if childRst == def.Running {
		return def.Running
	}

	//结果 成功
	if childRst == def.Successed {
		return def.Failed
	}

	//结果 失败
	if childRst == def.Failed {
		return def.Successed
	}

	return def.Failed
}
