package main

import (
	"fmt"
	"testcode/behavior_tree"
	"testcode/behavior_tree/node"
	"testcode/behavior_tree/node/action_node"
	"testcode/behavior_tree/node/logic_node"
	"time"
)

type TestI interface {
	Process()
	F1()
}

type BaseNode struct {
}

func (this *BaseNode) Process() {
	this.F1()
}

func (this *BaseNode) F1() {
	fmt.Printf("this is base node f1.\n")
}

type HNode struct {
	BaseNode
}

func (this *HNode) F1() {
	fmt.Printf("this is h node f1.\n")
}

//
//func (this *HNode) Process() {
//	fmt.Printf("this is h node.\n")
//}

func main2() {

	//snode := logic_node.NewBtNodeSequence("s", 0)
	//var baseNode node.BtNodeInterf
	//baseNode = snode
	//baseNode.Process()

	h := HNode{}
	h.BaseNode = BaseNode{}

	var testi TestI
	testi = &h

	testi.Process()
	testi.F1()
}

func main() {

	var s node.BtNodeInterf
	s = logic_node.NewBtNodeSequence("s", 0)
	var root node.BtNodeLogicInterf
	root = s.(node.BtNodeLogicInterf)

	a1 := action_node.NewBtNodeAction1("action1", 0)
	a2 := action_node.NewBtNodeAction2("action2", 0)

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
