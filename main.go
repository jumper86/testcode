package main

import (
	node2 "test/behavior_tree/node"
	"test/behavior_tree/node/logic_node"
)

func main() {

	node := logic_node.NewBtNodeSequence("s", 0)
	var baseNode node2.BtNodeInterf
	baseNode = &node
	baseNode.Evaluate()
}
