package behavior_tree

type BtDecoratorNode struct {
	BtNode
	child BtNodeInterf //所有子节点

}

func (this *BtDecoratorNode) GetChild() BtNodeInterf {
	return this.child
}

func (this *BtDecoratorNode) SetChild(bn BtNodeInterf) {
	if this.child == nil {
		this.child = bn
	}

	this.Reset()
	return
}

func (this *BtDecoratorNode) CleanChild() {
	this.Reset()
	this.child = nil
	return
}

//准入失败时，执行成功，执行失败时调用
func (this *BtDecoratorNode) Reset() {
	this.status = Ready
	this.child.Reset()
}
