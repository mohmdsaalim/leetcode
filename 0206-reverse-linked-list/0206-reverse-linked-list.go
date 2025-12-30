/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next   // save next
		current.Next = prev    // reverse link
		prev = current         // move prev
		current = next         // move current
	}

	return prev
}