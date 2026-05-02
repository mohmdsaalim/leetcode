/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
    if head == nil{
        return nil
    }

    var stack []*ListNode
    var current = head

    for current != nil{
        for len(stack) > 0 && current.Val > stack[len(stack)-1].Val{
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, current)
        current = current.Next
    }

    var newHead *ListNode

    for i := len(stack) -1; i >= 0; i--{
        stack[i].Next = newHead
        newHead = stack[i]
    }
    return newHead
}