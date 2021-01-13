package action_node

import (
	"fmt"
	"testcode/behavior_tree/def"
	"testcode/behavior_tree/node"
)

type BtNodeAction2 struct {
	node.BtNode
}

func NewBtNodeAction2(name string, interval int64) *BtNodeAction2 {
	var btns BtNodeAction2
	btns.BtNode = node.NewBtNode(name, interval)
	btns.SetTypes(def.Action1)
	return &btns
}

func (this *BtNodeAction2) Evaluate() bool {
	return true
}

func (this *BtNodeAction2) Tick() def.BtnResult {

	fmt.Printf("this is action 2 .\n")
	return def.Successed
}
