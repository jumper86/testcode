package decorator_node

import (
	"testcode/behavior_tree/def"
	"testcode/behavior_tree/node"
)

type BtNodeInverter struct {
	BtDecoratorNode
	child node.BtNodeInterf //所有子节点
}

func NewBtNodeInverter(name string, interval int64) *BtNodeInverter {
	var btns BtNodeInverter
	btns.BtNode = node.NewBtNode(name, interval)
	btns.SetTypes(def.DecoratorInverterNode)
	return &btns
}

func (this *BtNodeInverter) Evaluate() bool {
	return node.Evaluate(this.child)
}

func (this *BtNodeInverter) Tick() def.BtnResult {

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
