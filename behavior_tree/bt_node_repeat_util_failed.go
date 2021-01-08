package behavior_tree

type BtNodeRepeatUtilFail struct {
	BtNode
	child BtNodeInterf //所有子节点
}

func (this *BtNodeRepeatUtilFail) Evaluate() bool {
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
func (this *BtNodeRepeatUtilFail) DoEvaluate() bool {
	return this.child.Evaluate()
}

func (this *BtNodeRepeatUtilFail) Tick() BtnResult {

	childRst := this.child.Tick()
	//结果 运行中
	if childRst == Running {
		return Running
	}

	//todo: 此处存在问题
	//比如该节点下方是 selector 那么就会导致在完成一次之后，第二次执行的时候 selector.DoEvaluate 没有被调用
	//是否需要添加 Process 函数，然后将所有tick 函数中的 子节点.tick换成 子节点.process
	if childRst == Successed {
		this.Reset()
		return Running
	}

	this.Reset()
	return Failed
}

//准入失败时，执行成功，执行失败时调用
func (this *BtNodeRepeatUtilFail) Reset() {
	this.status = Ready
	this.child.Reset()
}

func (this *BtNodeRepeatUtilFail) AddChild(bn BtNodeInterf) {
	this.child = bn
	this.Reset()
	return
}

func (this *BtNodeRepeatUtilFail) RemoveChild(bn BtNodeInterf) {
	this.Reset()
	this.child = nil
	return
}
