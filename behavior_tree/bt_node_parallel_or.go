package behavior_tree

type BtNodeParallelOr struct {
	BtNodeCompose
	result []BtnResult //每个子节点对应执行结果

}

func NewBtNodeParallelOr(name string, interval int64) BtNodeParallelOr {
	var btns BtNodeParallelOr
	btns.BtNodeCompose = NewBtNodeCompose(name, interval)
	btns.result = make([]BtnResult, 0)
	btns.types = ComposeParallelOrNode

	return btns
}

//Evaluate 只在开始执行该节点时调用一次
func (this *BtNodeParallelOr) DoEvaluate() bool {
	if len(this.children) == 0 {
		return false
	}

	for _, child := range this.children {
		if child.Evaluate() {
			return true
		}
	}

	return false
}

//一个子节点失败就失败
//一个子节点成功就成功
func (this *BtNodeParallelOr) Tick() BtnResult {

	for _, child := range this.children {
		localRst := child.Tick()
		if localRst == Failed {
			this.Reset()
			return Failed
		}
		if localRst == Successed {
			this.Reset()
			return Successed
		}
	}

	return Running
}

//准入失败时，执行成功，执行失败时调用
func (this *BtNodeParallelOr) Reset() {
	this.BtNodeCompose.Reset()
	for i := range this.result {
		this.result[i] = Ready
	}
	for _, child := range this.children {
		child.Reset()
	}
}

func (this *BtNodeParallelOr) AddChild(bn BtNodeInterf) {
	if this.children == nil {
		this.children = make([]BtNodeInterf, 0)
		this.result = make([]BtnResult, 0)
	}
	if bn != nil {
		this.children = append(this.children, bn)
		this.result = append(this.result, Ready)
	}

	this.Reset()
	return
}

func (this *BtNodeParallelOr) RemoveChild(bn BtNodeInterf) {
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
		this.result = append(this.result[:objIdx], this.result[objIdx+1:]...)

	}

	this.Reset()
	return
}
