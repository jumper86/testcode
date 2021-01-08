package behavior_tree

type BtNodeInverter struct {
	BtNode
	child BtNodeInterf //所有子节点
}

func (this *BtNodeInverter) Evaluate() bool {
	//note: 保证只在第一次执行组合节点的时候，进行一次准入检查，即调用 Evaluate
	if this.status != Ready {
		return true
	}
	if this.activated && this.CheckTimer() && this.DoEvaluate() {
		this.status = Running
		return true
	}
	return false
}

//Evaluate 只在开始执行该节点时调用一次
func (this *BtNodeInverter) DoEvaluate() bool {
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

//准入失败时，执行成功，执行失败时调用
func (this *BtNodeInverter) Reset() {
	this.status = Ready
	this.child.Reset()
}

func (this *BtNodeInverter) AddChild(bn BtNodeInterf) {
	this.child = bn
	this.Reset()
	return
}

func (this *BtNodeInverter) RemoveChild(bn BtNodeInterf) {
	this.Reset()
	this.child = nil
	return
}

//
//func (this *BtNodeInverter) Process() BtnResult {
//	if this.Evaluate() {
//		return this.Tick()
//	}
//	return Failed
//}
