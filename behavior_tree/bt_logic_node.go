package behavior_tree

type BtLogicNode struct {
	BtNode
	children []BtNodeInterf //所有子节点

}

func (this *BtLogicNode) GetChildren() []BtNodeInterf {
	return this.children
}

func (this *BtLogicNode) AddChild(bn BtNodeInterf) {
	if this.children == nil {
		this.children = make([]BtNodeInterf, 0)
	}
	if bn != nil {
		this.children = append(this.children, bn)
	}

	this.Reset()
	return
}

func (this *BtLogicNode) RemoveChild(bn BtNodeInterf) {
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
	}

	this.Reset()
	return
}
