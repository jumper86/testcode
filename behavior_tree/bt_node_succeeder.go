package behavior_tree

type BtNodeSucceeder struct {
	BtDecoratorNode
	child BtNodeInterf //所有子节点
}

func NewBtNodeSucceeder(name string, interval int64) *BtNodeSucceeder {
	var btns BtNodeSucceeder
	btns.BtNode = NewBtNode(name, interval)
	btns.types = DecoratorSucceederNode
	btns.evaluate = btns.doEvaluate
	return &btns
}

func (this *BtNodeSucceeder) doEvaluate() bool {
	return this.child.Evaluate()
}

func (this *BtNodeSucceeder) Tick() BtnResult {

	childRst := this.child.Tick()
	//结果 运行中
	if childRst == Running {
		return Running
	}

	this.Reset()
	return Successed
}
