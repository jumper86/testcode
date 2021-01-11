package behavior_tree

import "testcode/behavior_tree/node"

type BehaviorTree struct {
	root node.BtNodeInterf
}

//每个帧调用该函数
func (this *BehaviorTree) Process() {
	if this.root.Evaluate() {
		this.root.Tick()
	}
}
