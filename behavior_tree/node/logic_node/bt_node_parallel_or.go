package logic_node

import (
	"testcode/behavior_tree/def"
	"testcode/behavior_tree/node"
)

//note:
//	检查方式：子节点中有一个满足，则检查通过
//	执行方式：每次Tick依次执行所有子节点的Process，只要有一个成功或者失败就立即结束
//	返回方式：
//		所有子节点返回running，则返回running
//		子节点返回failed，则返回failed
//		子节点返回successed，则返回successed

//note:
//	这里有一个比较特殊的情况，比如先后加入两个节点a(1000ms) b(100ms)
//	执行的时候，每100ms的时候， Evaluate 通过，然后执行第一个子节点a，成功返回
//	这就造成了，a节点每100ms执行一次的结果
//	这是因为在函数 Process 中对于行动节点取消判断 Evaluate 造成的。
//	但是这是正常的，因为 ComposeParallelOrNode 逻辑就是检查有一个通过就行，执行的时候都可以尝试执行，没毛病。
//	只是这个奇怪的结果稍加解释，这并非bug

type BtNodeParallelOr struct {
	BtLogicNode
	result []def.BtnResult //每个子节点对应执行结果

}

func NewBtNodeParallelOr(name string, interval int64) *BtNodeParallelOr {
	var btns BtNodeParallelOr
	btns.BtNode = node.NewBtNode(name, interval)
	btns.children = make([]node.BtNodeInterf, 0)
	btns.result = make([]def.BtnResult, 0)
	btns.SetTypes(def.ComposeParallelOrNode)

	return &btns
}

//Evaluate 只在开始执行该节点时调用一次
func (this *BtNodeParallelOr) Evaluate() bool {
	if len(this.children) == 0 {
		return false
	}

	for _, child := range this.children {
		if node.Evaluate(child) {
			return true
		}
	}

	return false
}

//一个子节点失败就失败
//一个子节点成功就成功
func (this *BtNodeParallelOr) Tick() def.BtnResult {

	for _, child := range this.children {
		localRst := node.Process(child)
		if localRst == def.Failed {
			return def.Failed
		}
		if localRst == def.Successed {
			return def.Successed
		}
	}

	return def.Running
}

func (this *BtNodeParallelOr) AddChild(bn node.BtNodeInterf) {
	if this.children == nil {
		this.children = make([]node.BtNodeInterf, 0)
		this.result = make([]def.BtnResult, 0)
	}
	if bn != nil {
		this.children = append(this.children, bn)
		this.result = append(this.result, def.None)
	}

	this.Reset()
	return
}

func (this *BtNodeParallelOr) RemoveChild(bn node.BtNodeInterf) {
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

//准入失败时，执行成功，执行失败时调用
func (this *BtNodeParallelOr) Reset() {
	this.SetStatus(def.Ready)
	for i := range this.result {
		this.result[i] = def.None
	}
	for _, child := range this.children {
		child.Reset()
	}
}
