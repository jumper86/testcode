package decorator_node

import (
	"testcode/behavior_tree/def"
	"testcode/behavior_tree/node"
)

type BtNodeSucceeder struct {
	BtDecoratorNode
	child node.BtNodeInterf //所有子节点
}

func NewBtNodeSucceeder(name string, interval int64) *BtNodeSucceeder {
	var btns BtNodeSucceeder
	btns.BtNode = node.NewBtNode(name, interval)
	btns.SetTypes(def.DecoratorSucceederNode)
	return &btns
}

func (this *BtNodeSucceeder) Evaluate() bool {
	return this.child.Evaluate()
}

func (this *BtNodeSucceeder) Tick() def.BtnResult {

	childRst := node.Process(this.child)
	//结果 运行中
	if childRst == def.Running {
		return def.Running
	}

	return def.Successed
}
