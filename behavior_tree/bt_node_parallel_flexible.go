package behavior_tree

type BtNodeParallelFlexible struct {
	BtNodeCompose
	evaluate []bool //每个子节点对应对应准入

}

func NewBtNodeParallelFlexible(name string, interval int64) BtNodeParallelFlexible {
	var btns BtNodeParallelFlexible
	btns.BtNodeCompose = NewBtNodeCompose(name, interval)
	btns.evaluate = make([]bool, 0)
	btns.types = ComposeParallelFlexibleNode

	return btns
}

//Evaluate 只在开始执行该节点时调用一次
func (this *BtNodeParallelFlexible) DoEvaluate() bool {
	if len(this.children) == 0 {
		return false
	}

	evaluateCnt := 0
	for i, child := range this.children {
		if child.Evaluate() {
			evaluateCnt++
			this.evaluate[i] = true
		}
	}

	if evaluateCnt > 0 {
		return true
	}
	return false
}

//一个子节点失败就失败
//所有子节点成功才成功
//一个子节点运行就返回运行
func (this *BtNodeParallelFlexible) Tick() BtnResult {

	//寻找处于running的子节点
	toTick := make([]int, 0)
	for i, eva := range this.evaluate {
		if eva {
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
func (this *BtNodeParallelFlexible) Reset() {
	this.BtNodeCompose.Reset()
	for i := range this.evaluate {
		this.evaluate[i] = false
	}
	for _, child := range this.children {
		child.Reset()
	}
}

func (this *BtNodeParallelFlexible) AddChild(bn BtNodeInterf) {
	if this.children == nil {
		this.children = make([]BtNodeInterf, 0)
		this.evaluate = make([]bool, 0)
	}
	if bn != nil {
		this.children = append(this.children, bn)
		this.evaluate = append(this.evaluate, false)
	}

	this.Reset()
	return
}

func (this *BtNodeParallelFlexible) RemoveChild(bn BtNodeInterf) {
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
		this.evaluate = append(this.evaluate[:objIdx], this.evaluate[objIdx+1:]...)

	}

	this.Reset()
	return
}
