package logic_node

import (
	"testcode/behavior_tree/def"
	"testcode/behavior_tree/node"
)

type BtNodeParallelFlexible struct {
	BtLogicNode
	evaluateRst []bool //每个子节点对应对应准入

}

func NewBtNodeParallelFlexible(name string, interval int64) *BtNodeParallelFlexible {
	var btns BtNodeParallelFlexible
	btns.BtNode = node.NewBtNode(name, interval)
	btns.children = make([]node.BtNodeInterf, 0)
	btns.evaluateRst = make([]bool, 0)
	btns.SetTypes(def.ComposeParallelFlexibleNode)
	btns.SetEvaluate(btns.doEvaluate)

	return &btns
}

func (this *BtNodeParallelFlexible) doEvaluate() bool {
	if len(this.children) == 0 {
		return false
	}

	evaluateCnt := 0
	for i, child := range this.children {
		if child.Evaluate() {
			evaluateCnt++
			this.evaluateRst[i] = true
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
func (this *BtNodeParallelFlexible) Tick() def.BtnResult {

	//寻找处于running的子节点
	toTick := make([]int, 0)
	for i, eva := range this.evaluateRst {
		if eva {
			toTick = append(toTick, i)
		}
	}
	runningCnt := len(toTick)

	//执行每个running子节点的tick
	for _, runningIdx := range toTick {
		localRst := this.children[runningIdx].Process()

		if localRst == def.Failed {
			return def.Failed
		}

		if localRst == def.Successed {
			runningCnt--
		}
	}

	if runningCnt == 0 {
		return def.Successed
	}

	return def.Running
}

func (this *BtNodeParallelFlexible) AddChild(bn node.BtNodeInterf) {
	if this.children == nil {
		this.children = make([]node.BtNodeInterf, 0)
		this.evaluateRst = make([]bool, 0)
	}
	if bn != nil {
		this.children = append(this.children, bn)
		this.evaluateRst = append(this.evaluateRst, false)
	}

	this.Reset()
	return
}

func (this *BtNodeParallelFlexible) RemoveChild(bn node.BtNodeInterf) {
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
		this.evaluateRst = append(this.evaluateRst[:objIdx], this.evaluateRst[objIdx+1:]...)

	}

	this.Reset()
	return
}

//准入失败时，执行成功，执行失败时调用
func (this *BtNodeParallelFlexible) Reset() {
	this.SetStatus(def.Ready)
	for i := range this.evaluateRst {
		this.evaluateRst[i] = false
	}
	for _, child := range this.children {
		child.Reset()
	}
}

func (this *BtNodeParallelFlexible) Process() def.BtnResult {
	if !this.Evaluate() {
		return def.Failed
	}
	if this.GetStatus() != def.Run {
		this.SetStatus(def.Run)
	}

	tmpRst := this.Tick()
	if tmpRst != def.Running {
		this.Reset()
	}
	return tmpRst
}
