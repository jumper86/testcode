package behavior_tree

type BtNodeCompose struct {
	BtNode
	children []BtNodeInterf //所有子节点
}

func NewBtNodeCompose(name string, interval int64) BtNodeCompose {
	var btnc BtNodeCompose
	btnc.BtNode = NewBtNode(name, interval)
	btnc.children = make([]BtNodeInterf, 0)
	return btnc
}

func (this *BtNodeCompose) Evaluate() bool {
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

func (this *BtNodeCompose) Reset() {
	this.status = Ready
}

func (this *BtNodeCompose) AddChild(bn BtNodeInterf) {
	if this.children == nil {
		this.children = make([]BtNodeInterf, 0)
	}
	if bn != nil {
		this.children = append(this.children, bn)
	}

	this.Reset()
	return
}

func (this *BtNodeCompose) RemoveChild(bn BtNodeInterf) {
	objId := bn.GetId()
	objIdx := -1
	for idx, child := range this.children {
		if child.GetId() == objId {
			objIdx = idx
			break
		}
	}
	if objIdx != -1 {
		this.children = append(this.children[:objIdx], this.children[objIdx+1:]...)
	}

	this.Reset()
	return
}
