package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp, odd, even := head, head, head.Next
	head = even

	for odd != nil && even != nil {
		odd.Next = even.Next
		even.Next = odd
		tmp = odd

		odd = odd.Next
		if odd != nil {
			even = odd.Next
			if even != nil {
				tmp.Next = even
			}
		}
	}
	return head
}

func main() {
	PrintList(swapPairs(Num2List(1, 2, 3, 4, 5, 6)))
	PrintList(swapPairs(nil))
}
