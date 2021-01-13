package main

import (
	"testcode/behavior_tree/node"
	"testcode/behavior_tree/node/logic_node"
)

func main() {

	snode := logic_node.NewBtNodeSequence("s", 0)
	var baseNode node.BtNodeInterf
	baseNode = snode
	baseNode.Process()
}
