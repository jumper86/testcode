package decorator_node

import (
	"test/behavior_tree/def"
	"test/behavior_tree/node"
)

type BtNodeRepeatUtilFail struct {
	BtDecoratorNode
	child node.BtNodeInterf //所有子节点
}

func NewBtNodeRepeatUtilFail(name string, interval int64) *BtNodeRepeatUtilFail {
	var btns BtNodeRepeatUtilFail
	btns.BtNode = node.NewBtNode(name, interval)
	btns.SetTypes(def.DecoratorRepeatUtilFailedNode)
	btns.SetEvaluate(btns.doEvaluate)
	return &btns
}

func (this *BtNodeRepeatUtilFail) doEvaluate() bool {
	return this.child.Evaluate()
}

func (this *BtNodeRepeatUtilFail) Tick() def.BtnResult {

	childRst := this.child.Tick()
	//结果 运行中
	if childRst == def.Running {
		return def.Running
	}

	//todo: 此处存在问题
	//比如该节点下方是 selector 那么就会导致在完成一次之后，第二次执行的时候 selector.DoEvaluate 没有被调用
	//是否需要添加 Process 函数，然后将所有tick 函数中的 子节点.tick换成 子节点.process
	if childRst == def.Successed {
		this.Reset()
		return def.Running
	}

	this.Reset()
	return def.Failed
}
