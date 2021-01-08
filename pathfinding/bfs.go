package main

import "fmt"

type Node struct {
	Name       string
	LeftChild  *Node
	RightChild *Node
}

type Tree struct {
	Root *Node //根节点
}

func composeTree() *Tree {
	h := &Node{
		Name:       "h",
		LeftChild:  nil,
		RightChild: nil,
	}
	f := &Node{
		Name:       "f",
		LeftChild:  nil,
		RightChild: nil,
	}
	g := &Node{
		Name:       "g",
		LeftChild:  nil,
		RightChild: nil,
	}

	d := &Node{
		Name:       "d",
		LeftChild:  nil,
		RightChild: nil,
	}

	e := &Node{
		Name:       "e",
		LeftChild:  h,
		RightChild: nil,
	}

	b := &Node{
		Name:       "b",
		LeftChild:  d,
		RightChild: e,
	}

	c := &Node{
		Name:       "c",
		LeftChild:  f,
		RightChild: g,
	}

	a := &Node{
		Name:       "a",
		LeftChild:  b,
		RightChild: c,
	}

	return &Tree{Root: a}

}

func findPath(tree *Tree, exitName string) *Node {
	r := tree.Root
	if r.Name == exitName {
		return r
	}

	tmp := make([]*Node, 0)
	if r.LeftChild != nil {
		tmp = append(tmp, r.LeftChild)
	}
	if r.RightChild != nil {
		tmp = append(tmp, r.RightChild)
	}

	var currentNode *Node
	for {
		if len(tmp) == 0 {
			currentNode = nil
			break
		}
		currentNode = tmp[0]
		tmp = tmp[1:]

		if currentNode.Name == exitName {
			break
		}

		if currentNode.LeftChild != nil {
			tmp = append(tmp, currentNode.LeftChild)
		}
		if currentNode.RightChild != nil {
			tmp = append(tmp, currentNode.RightChild)
		}
	}

	return currentNode
}

func main() {
	t := composeTree()
	rn := findPath(t, "f")
	if rn != nil {
		fmt.Printf("find %s\n", rn.Name)
	}
}
