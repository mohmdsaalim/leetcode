/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
    if root == nil {
        return false
    }
    
    return match(head, root) || 
           isSubPath(head, root.Left) || 
           isSubPath(head, root.Right)
}

func match(head *ListNode, root *TreeNode) bool {
    if head == nil {
        return true
    }
    if root == nil {
        return false
    }
    
    if head.Val != root.Val {
        return false
    }
    
    return match(head.Next, root.Left) || 
           match(head.Next, root.Right)
}


