package decorator_node

import (
	"test/behavior_tree/def"
	"test/behavior_tree/node"
)

type BtNodeSucceeder struct {
	BtDecoratorNode
	child node.BtNodeInterf //所有子节点
}

func NewBtNodeSucceeder(name string, interval int64) *BtNodeSucceeder {
	var btns BtNodeSucceeder
	btns.BtNode = node.NewBtNode(name, interval)
	btns.SetTypes(def.DecoratorSucceederNode)
	btns.SetEvaluate(btns.doEvaluate)
	return &btns
}

func (this *BtNodeSucceeder) doEvaluate() bool {
	return this.child.Evaluate()
}

func (this *BtNodeSucceeder) Tick() def.BtnResult {

	childRst := this.child.Tick()
	//结果 运行中
	if childRst == def.Running {
		return def.Running
	}

	this.Reset()
	return def.Successed
}
