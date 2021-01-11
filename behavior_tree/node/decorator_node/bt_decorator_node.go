package decorator_node

import (
	"test/behavior_tree/def"
	"test/behavior_tree/node"
)

type BtDecoratorNode struct {
	node.BtNode
	child node.BtNodeInterf //所有子节点

}

func (this *BtDecoratorNode) GetChild() node.BtNodeInterf {
	return this.child
}

func (this *BtDecoratorNode) SetChild(bn node.BtNodeInterf) {
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
	this.SetStatus(def.Ready)
	this.child.Reset()
}
