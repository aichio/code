package main

import (
	tree "code/mooc/treeentry_embedded"
	"fmt"
)

type myTreeNode struct {
	*tree.Node
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("This method is shadowed.")
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}

func main() {
	var root tree.Node
	fmt.Println(root)

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Traverse()
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
