/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    if head.Next.Next == nil{
        if head.Val > head.Next.Val{
            head.Val, head.Next.Val = head.Next.Val, head.Val
        }
        return head
    }
    mid := getMid(head)
    left := head
    right := mid.Next
    mid.Next = nil
    return merge(sortList(left), sortList(right))
}

func merge(left, right *ListNode) *ListNode{
    stub := &ListNode{}
    cur := stub
    for left != nil && right != nil{
        if left.Val < right.Val{
            cur.Next = left
            left = left.Next 
        } else {
            cur.Next = right
            right = right.Next 
        }
        cur = cur.Next
    }
    if left != nil {
        cur.Next = left
    } 
    if right != nil {
        cur.Next = right
    }
    return stub.Next
}

func getMid(head *ListNode) *ListNode{
    slow, fast := head, head
    for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil{
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}