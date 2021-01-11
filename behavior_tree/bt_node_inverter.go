package behavior_tree

type BtNodeInverter struct {
	BtDecoratorNode
	child BtNodeInterf //所有子节点
}

func NewBtNodeInverter(name string, interval int64) *BtNodeInverter {
	var btns BtNodeInverter
	btns.BtNode = NewBtNode(name, interval)
	btns.types = DecoratorInverterNode
	btns.evaluate = btns.doEvaluate
	return &btns
}

func (this *BtNodeInverter) doEvaluate() bool {
	return this.child.Evaluate()
}

func (this *BtNodeInverter) Tick() BtnResult {

	childRst := this.child.Tick()
	//结果 运行中
	if childRst == Running {
		return Running
	}

	//结果 成功
	if childRst == Successed {
		this.Reset()
		return Failed
	}

	//结果 失败
	if childRst == Failed {
		this.Reset()
		return Successed
	}

	return Failed
}
