package behavior_tree

import "testcode/behavior_tree/node"

type BehaviorTree struct {
	root node.BtNodeInterf
}

func NewBehaviorTree(root node.BtNodeInterf) *BehaviorTree {
	return &BehaviorTree{root: root}
}

//每个帧调用该函数
func (this *BehaviorTree) Process() {
	node.Process(this.root)
}
