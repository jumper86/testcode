package behavior_tree

type BtNodeParallelAnd struct {
	BtNodeCompose
	result []BtnResult //每个子节点对应执行结果

}

func NewBtNodeParallelAnd(name string, interval int64) BtNodeParallelAnd {
	var btns BtNodeParallelAnd
	btns.BtNodeCompose = NewBtNodeCompose(name, interval)
	btns.result = make([]BtnResult, 0)
	return btns
}

//Evaluate 只在开始执行该节点时调用一次
func (this *BtNodeParallelAnd) DoEvaluate() bool {
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

//一个子节点失败就失败
//所有子节点成功才成功
//一个子节点运行就返回运行
func (this *BtNodeParallelAnd) Tick() BtnResult {

	//寻找处于running的子节点
	toTick := make([]int, 0)
	for i, result := range this.result {
		if result == Ready || result == Running {
			toTick = append(toTick, i)
		}
	}
	runningCnt := len(toTick)

	//执行每个running子节点的tick
	for _, runningIdx := range toTick {
		localRst := this.children[runningIdx].Tick()

		if localRst == Failed {
			this.Reset()
			return Failed
		}

		if localRst == Successed {
			runningCnt--
		}
	}

	if runningCnt == 0 {
		this.Reset()
		return Successed
	}

	return Running
}

//准入失败时，执行成功，执行失败时调用
func (this *BtNodeParallelAnd) Reset() {
	this.BtNodeCompose.Reset()
	for i := range this.result {
		this.result[i] = Ready
	}
	for _, child := range this.children {
		child.Reset()
	}
}

func (this *BtNodeParallelAnd) AddChild(bn BtNodeInterf) {
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

func (this *BtNodeParallelAnd) RemoveChild(bn BtNodeInterf) {
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
