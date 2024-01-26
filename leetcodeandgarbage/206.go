

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {

// 		return head
// 	}
// 	newHead := reverseList(head.Next) // Рекурсивно получаем новый head

// 	// Устанавливаем следующий указатель у текущего head на nil
// 	head.Next.Next = head
// 	head.Next = nil

// 	return newHead // Возвращаем новый head
// }
// func main() {
// 	n1 := &ListNode{Val: 1}
// 	n2 := &ListNode{Val: 2}
// 	n1.Next = n2
// 	n3 := &ListNode{Val: 3}
// 	n2.Next = n3
// 	n4 := &ListNode{Val: 4}
// 	n3.Next = n4
// 	n5 := &ListNode{Val: 5}
// 	n4.Next = n5
// 	cur := n1
// 	for cur != nil {
// 		splitter := ""
// 		if cur != n1 {
// 			splitter = ">"
// 		}
// 		fmt.Printf("%s%d ", splitter, cur.Val)
// 		cur = cur.Next
// 	}
// }



package main

// Given the head of a singly linked list, 
// reverse the list, and return the reversed list.
// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
type ListNode struct {
	Val  int
	Next *ListNode
}

func TestList(head *ListNode) *ListNode {

	var tmp *ListNode = nil
	var prev *ListNode = nil
	var curr *ListNode = head

	for curr != nil { // 1
		tmp = curr.Next // 2

		curr.Next = prev // 0
		prev = curr      // 1
		curr = tmp       // 2
	}

	return prev
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n1.Next = n2
	n3 := &ListNode{Val: 3}
	n2.Next = n3
	TestList(n1)
}
