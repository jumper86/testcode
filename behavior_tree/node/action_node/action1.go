package action_node

import (
	"fmt"
	"testcode/behavior_tree/def"
	"testcode/behavior_tree/node"
)

type BtNodeAction1 struct {
	node.BtNode
}

func NewBtNodeAction1(name string, interval int64) *BtNodeAction1 {
	var btns BtNodeAction1
	btns.BtNode = node.NewBtNode(name, interval)
	btns.SetTypes(def.Action1)
	return &btns
}

func (this *BtNodeAction1) Evaluate() bool {
	return true
}

func (this *BtNodeAction1) Tick() def.BtnResult {

	fmt.Printf("this is action 1 .\n")
	return def.Successed
}
