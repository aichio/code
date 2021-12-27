package tree

import "fmt"

type Node struct {
	value       int
	left, right *Node
}

func CreateNode(value int) *Node {
	return &Node{value: value}
}

func (node Node) Print() {
	fmt.Println(node.value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil" + "node. Ignored.")
	}
	node.value = value
}
