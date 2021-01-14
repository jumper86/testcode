package main

import (
	"testcode/behavior_tree"
	"testcode/behavior_tree/node"
	"testcode/behavior_tree/node/action_node"
	"testcode/behavior_tree/node/logic_node"
	"time"
)

func main() {

	var s node.BtNodeInterf
	s = logic_node.NewBtNodeSelector("s", 0)
	var root node.BtNodeLogicInterf
	root = s.(node.BtNodeLogicInterf)

	a1 := action_node.NewBtNodeAction1("action1", 1000)
	a2 := action_node.NewBtNodeAction2("action2", 100)

	root.AddChild(a1)
	root.AddChild(a2)

	tree := behavior_tree.NewBehaviorTree(root)
	ticker := time.NewTicker(time.Second / time.Duration(60))
	for {
		select {
		case <-ticker.C:
			tree.Process()
		}
	}
}
