package main

import "fmt"

type Node struct {
	Value               int
	left, right, parent *Node
}

func NewNode(value int, left *Node, right *Node) *Node {
	n := &Node{Value: value, left: left, right: right}
	left.parent = n
	right.parent = n
	return n
}

func NewTerminalNode(value int) *Node {
	return &Node{Value: value}
}

type InorderIterator struct {
	Current       *Node
	root          *Node
	returnedStart bool
}

func NewInorderIterator(root *Node) *InorderIterator {
	i := &InorderIterator{root, root, false}
	for i.Current.left != nil {
		i.Current = i.Current.left
	}
	return i
}

func main() {
	//	1
	// / \
	//2   3

	// inorder: 213
	//preorder: 123
	//postorder: 231

	root := NewNode(1, NewTerminalNode(2), NewTerminalNode(3))
	fmt.Println(root)

}
