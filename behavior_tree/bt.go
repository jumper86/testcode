package behavior_tree

type BehaviorTree struct {
	root BtNodeInterf
}

//每个帧调用该函数
func (this *BehaviorTree) Process() {
	if this.root.Evaluate() {
		this.root.Tick()
	}
}
