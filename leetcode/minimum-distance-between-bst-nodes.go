package main

var prevNode *TreeNode = nil

func minDiffInBST(root *TreeNode) int {
	minDiff := 1000000
	node := root
	if node == nil {
		return 1000000
	}
	if node.Left != nil {
		diff := minDiffInBST(node.Left)
		if diff < minDiff {
			minDiff = diff
		}
	}
	if prevNode != nil {
		diff := prevNode.Val - node.Val
		if diff < 0 {
			diff = -diff
		}
		if diff < minDiff {
			minDiff = diff
		}
	}
	prevNode = node
	if node.Right != nil {
		diff := minDiffInBST(node.Right)
		if diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}

func main() {

}
