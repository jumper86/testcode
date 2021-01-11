package behavior_tree

type BtNodeSelector struct {
	BtNode
	children    []BtNodeInterf //所有子节点
	activeIdx   int
	activeChild BtNodeInterf //当前执行子节点
}

func NewBtNodeSelector(name string, interval int64) *BtNodeSelector {
	var btns BtNodeSelector
	btns.BtNode = NewBtNode(name, interval)
	btns.children = make([]BtNodeInterf, 0)
	btns.activeIdx = -1
	btns.activeChild = nil
	btns.types = ComposeSelectorNode
	btns.evaluate = btns.doEvaluate

	return &btns
}

func (this *BtNodeSelector) doEvaluate() bool {
	if len(this.children) == 0 {
		return false
	}

	for i, child := range this.children {
		if child.Evaluate() {
			this.activeIdx = i
			this.activeChild = child
			return true
		}
	}

	return false
}

func (this *BtNodeSelector) Tick() BtnResult {

	if this.activeChild == nil {
		return Failed
	}

	childRst := this.activeChild.Tick()

	//运行中
	if childRst == Running {
		return Running
	}

	//成功
	if childRst == Successed {
		this.Reset()
		return Successed
	}

	//失败
	if childRst == Failed {
		//寻找下一个可以执行的子节点
		found := false
		for i := this.activeIdx; i < len(this.children); i++ {
			tmp := this.children[i]
			if tmp.Evaluate() {
				found = true
				this.activeIdx = i
				this.activeChild = tmp
				break
			}
		}

		if !found {
			this.Reset()
			return Failed
		} else {
			return Running
		}
	}

	return Failed
}

func (this *BtNodeSelector) Reset() {
	this.status = Ready
	this.activeIdx = -1
	this.activeChild = nil
	for _, child := range this.children {
		child.Reset()
	}
}
