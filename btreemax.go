package piscine

type TreeNode struct {
	Parent, Left, Right *TreeNode
	Data                string
}

func BTreeMax(root *TreeNode) *TreeNode {
	max := root
	
	for root.Right != nil {
		max = root.Right
		root = root.Right
	}
	return max

}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {

	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root

	} else {
		root.Right = BTreeInsertData(root.Right, data)
		root.Right.Parent = root
	}
	return root
}
