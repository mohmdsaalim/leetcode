/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
    if head == nil{
        return nil
    }

    valueSet := make(map[int]bool)
    for _, value := range nums{
        valueSet[value] = true
    }
    dummy := &ListNode{}
    dummy.Next = head
    previousNode := dummy
    currentNode := head

    for currentNode != nil{
        if valueSet[currentNode.Val] {
          previousNode.Next = currentNode.Next
        }else{
                 previousNode = currentNode
        }
        currentNode = currentNode.Next
    }
    return dummy.Next
}