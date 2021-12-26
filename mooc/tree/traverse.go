package tree

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.left.Traverse()
	node.Print()
	node.right.Traverse()
}
