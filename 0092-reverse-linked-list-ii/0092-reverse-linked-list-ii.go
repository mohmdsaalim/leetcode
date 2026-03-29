/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil || left == right {
        return head
    }

    dummy := &ListNode{Next: head}
    prev := dummy

    // Step 1: Move prev to node before 'left'
    for i := 1; i < left; i++ {
        prev = prev.Next
    }

    // Step 2: Start reversing
    curr := prev.Next
    var next *ListNode

    for i := 0; i < right-left; i++ {
        next = curr.Next
        curr.Next = next.Next
        next.Next = prev.Next
        prev.Next = next
    }

    return dummy.Next
}