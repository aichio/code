package main

import (
	"fmt"
	"mooc/tree"
)

type myTreeNode struct {
	*tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode.node == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	fmt.Println(root)

	root = tree.Node{value: 3}
	root.left = &tree.Node{}
	root.right = &tree.Node{5, nil, nil}
	root.right.left = new(tree.Node)
	root.left.right = CreateNode(2)
	root.traverse()
	// fmt.Println(root)
	// nodes := []tree.Node{
	// 	{value: 3},
	// 	{},
	// 	{5, nil, &root},
	// }
	// fmt.Println(nodes)
	// root.print()
	// root.setValue(100)
	// root.print()
	// var pRoot *tree.Node
	// pRoot.setValue(200)
	// pRoot = &root
	// pRoot.setValue(300)
	// pRoot.print()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

}
