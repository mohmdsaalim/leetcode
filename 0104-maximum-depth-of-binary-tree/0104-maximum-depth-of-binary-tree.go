/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil{
        return 0
    }

    right := maxDepth(root.Right)
    left := maxDepth(root.Left)

    return 1 + max(right, left)
}