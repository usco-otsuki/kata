class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def solve(self, root: TreeNode):
        if root is None:
            return 0, 0

        left_max_depth, left_max_diameter = self.solve(root.left)
        right_max_depth, right_max_diameter = self.solve(root.right)
        return max(left_max_depth, right_max_depth) + 1, max(left_max_depth + right_max_depth , left_max_diameter, right_max_diameter)

    def diameterOfBinaryTree(self, root: TreeNode) -> int:
        max_depth, max_diameter = self.solve(root)
        return max_diameter

if __name__ == "__main__":
    solver = Solution()

    left = TreeNode(4)
    right = TreeNode(3)
    node = TreeNode(2)
    node.left = left
    node.right = right
    left = node
    right = TreeNode(3)
    node = TreeNode(1)
    node.left = left
    node.right = right
    assert solver.diameterOfBinaryTree(node) == 3
