package decorator_node

import (
	"test/behavior_tree/def"
	"test/behavior_tree/node"
)

type BtNodeInverter struct {
	BtDecoratorNode
	child node.BtNodeInterf //所有子节点
}

func NewBtNodeInverter(name string, interval int64) *BtNodeInverter {
	var btns BtNodeInverter
	btns.BtNode = node.NewBtNode(name, interval)
	btns.SetTypes(def.DecoratorInverterNode)
	btns.SetEvaluate(btns.doEvaluate)
	return &btns
}

func (this *BtNodeInverter) doEvaluate() bool {
	return this.child.Evaluate()
}

func (this *BtNodeInverter) Tick() def.BtnResult {

	childRst := this.child.Tick()
	//结果 运行中
	if childRst == def.Running {
		return def.Running
	}

	//结果 成功
	if childRst == def.Successed {
		this.Reset()
		return def.Failed
	}

	//结果 失败
	if childRst == def.Failed {
		this.Reset()
		return def.Successed
	}

	return def.Failed
}
