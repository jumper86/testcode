package behavior_tree

type BtNodeSequence struct {
	BtNodeCompose
	activeIdx   int          //当前执行子节点idx
	activeChild BtNodeInterf //当前执行子节点
}

func NewBtNodeSequence(name string, interval int64) BtNodeSequence {
	var btns BtNodeSequence
	btns.BtNodeCompose = NewBtNodeCompose(name, interval)
	btns.activeIdx = -1
	btns.activeChild = nil
	return btns
}

//Evaluate 只在开始执行该节点时调用一次
func (this *BtNodeSequence) DoEvaluate() bool {
	if len(this.children) == 0 {
		return false
	}

	for _, child := range this.children {
		if !child.Evaluate() {
			return false
		}
	}

	return true
}

func (this *BtNodeSequence) Tick() BtnResult {
	if this.activeChild == nil {
		this.activeIdx = 0
		this.activeChild = this.children[0]
	}

	//执行逻辑
	childRst := this.activeChild.Tick()

	//结果 运行中
	if childRst == Running {
		return Running
	}

	//结果 成功
	if childRst == Successed {
		if this.activeIdx == len(this.children)-1 {
			this.Reset()
			return Successed
		} else {
			this.activeIdx++
			this.activeChild = this.children[this.activeIdx]
			return Running
		}
	}

	//结果 失败
	if childRst == Failed {
		this.Reset()
		return Failed
	}

	return Failed
}

//准入失败时，执行成功，执行失败时调用
func (this *BtNodeSequence) Reset() {
	this.BtNodeCompose.Reset()
	this.activeIdx = -1
	this.activeChild = nil
	for _, child := range this.children {
		child.Reset()
	}
}
