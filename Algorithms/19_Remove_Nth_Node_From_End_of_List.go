package main

import (
	"github.com/foresee88/leetcode_go/Algorithms/common"
)

/*

19. Remove Nth Node From End of List
Medium

1975

149

Favorite

Share
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = common.ListNode

func getListLen(l *ListNode) int {
	i := 0
	for p := l; p != nil; p = p.Next {
		i++
	}
	return i
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := getListLen(head)
	if l == 0 || n == 0 || l < n {
		return head
	}
	if l-n == 0 {
		head = head.Next
		return head
	}
	for p, i := head, 0; i <= l-n; p, i = p.Next, i+1 {

		if i == l-n-1 {
			p.Next = p.Next.Next
			// 如果n=1,此时p.Next已经是nil了，因此不break，下一次循环p.Next会报错
			break
		}

	}
	return head
}

func main() {
	common.PrintList(removeNthFromEnd(common.Num2List(1, 2, 3, 4, 5), 1))
	common.PrintList(removeNthFromEnd(common.Num2List(1, 2, 3, 4, 5), 0))
	common.PrintList(removeNthFromEnd(common.Num2List(1, 2, 3, 4, 5), 5))
	common.PrintList(removeNthFromEnd(common.Num2List(1, 2, 3, 4, 5), 2))
}
