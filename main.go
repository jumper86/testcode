package main

import "test/behavior_tree"

func main() {

	node := behavior_tree.NewBtNodeSequence("s", 0)
	var baseNode behavior_tree.BtNodeInterf
	baseNode = &node
	baseNode.Evaluate()
}
