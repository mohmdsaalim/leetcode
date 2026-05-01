/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
     dummy := list1
    str, fin := &ListNode{}, &ListNode{}
    for list1 != nil {
        if a == 1 {
            str = list1
        }
        if b == -1 {
            fin = list1
            break
        }
        list1 = list1.Next
        a--
        b--
    }
    str.Next = list2
    for list2.Next != nil {
        list2 = list2.Next
    }
    list2.Next = fin
    return dummy
}