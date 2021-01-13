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
	btns.SetEvaluate(btns.doEvaluate)
	return &btns
}

func (this *BtNodeSucceeder) doEvaluate() bool {
	return this.child.Evaluate()
}

func (this *BtNodeSucceeder) Tick() def.BtnResult {

	childRst := this.child.Process()
	//结果 运行中
	if childRst == def.Running {
		return def.Running
	}

	return def.Successed
}

func (this *BtNodeSucceeder) Process() def.BtnResult {
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
