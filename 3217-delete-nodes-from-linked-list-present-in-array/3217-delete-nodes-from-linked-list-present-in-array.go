/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func modifiedList(nums []int, head *ListNode) *ListNode {
//     if head == nil{
//         return nil
//     }

//     valueSet := make(map[int]bool)
//     for _, value := range nums{
//         valueSet[value] = true
//     }
//     dummy := &ListNode{}
//     dummy.Next = head
//     previousNode := dummy
//     currentNode := head

//     for currentNode != nil{
//         if valueSet[currentNode.Val] {
//           previousNode.Next = currentNode.Next
//         }else{
//                  previousNode = currentNode
//         }
//         currentNode = currentNode.Next
//     }
//     return dummy.Next
// }

func modifiedList(nums []int, head *ListNode) *ListNode {
    valueSet := make(map[int]struct{})
    for _, v := range nums {
        valueSet[v] = struct{}{}
    }

    dummy := &ListNode{Next: head}
    prev, curr := dummy, head

    for curr != nil {
        if _, exists := valueSet[curr.Val]; exists {
            prev.Next = curr.Next
        } else {
            prev = curr
        }
        curr = curr.Next
    }

    return dummy.Next
}