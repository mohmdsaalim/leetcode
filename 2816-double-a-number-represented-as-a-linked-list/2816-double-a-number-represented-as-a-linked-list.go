/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
    if head.Val > 4{
        newHead := &ListNode{Val: 0, Next: head}
        head = newHead
    }

    current := head

    for current != nil{
        current.Val = (current.Val * 2) % 10

        if current.Next != nil && current.Next.Val > 4{
            current.Val += 1
        }
        current = current.Next
    }
    return head

}