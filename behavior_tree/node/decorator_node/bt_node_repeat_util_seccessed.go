package decorator_node

import (
	"testcode/behavior_tree/def"
	"testcode/behavior_tree/node"
)

type BtNodeRepeatUtilSeccessed struct {
	BtDecoratorNode
	child node.BtNodeInterf //所有子节点
}

func NewBtNodeRepeatUtilSeccessed(name string, interval int64) *BtNodeRepeatUtilSeccessed {
	var btns BtNodeRepeatUtilSeccessed
	btns.BtNode = node.NewBtNode(name, interval)
	btns.SetTypes(def.DecoratorRepeatUtilFailedNode)
	btns.SetEvaluate(btns.doEvaluate)
	return &btns
}

func (this *BtNodeRepeatUtilSeccessed) doEvaluate() bool {
	return this.child.Evaluate()
}

func (this *BtNodeRepeatUtilSeccessed) Tick() def.BtnResult {

	childRst := this.child.Process()
	//结果 运行中
	if childRst == def.Running {
		return def.Running
	}

	//比如该节点下方是 selector 那么就会导致在完成一次之后，第二次执行的时候 需要再次调用 doEvaluate
	//因此这里需要调用 child.Reset
	if childRst == def.Failed {
		this.child.Reset()
		return def.Running
	}

	return def.Successed
}

func (this *BtNodeRepeatUtilSeccessed) Process() def.BtnResult {
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
