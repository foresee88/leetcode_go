package main

import (
	"github.com/foresee88/leetcode_go/Algorithms/common"
)

/*
25. Reverse Nodes in k-Group
Hard

1264

259

Favorite

Share
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:

Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5

Note:

Only constant extra memory is allowed.
You may not alter the values in the list's nodes, only nodes itself may be changed.
*/

/*
* Definition for singly-linked list.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = common.ListNode

func reverseList(l *ListNode) (head, end *ListNode) {
	var h, e *ListNode
	// 取到列表最后一个元素r，并获取列表长度k
	k := 0
	for p := l; p != nil; p = p.Next {
		k++
		h = p
	}
	for q := h; k > 0; q = q.Next {
		k--
		n := 0
		for p := l; p != nil; p = p.Next {
			if n == k {
				q.Next = p
				if n == 0 {
					p.Next = nil
					e = p
				}
				break
			}
			n++
		}
	}

	return h, e

}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	// 获取列表长度l,如果小于k，直接返回列表（题面排除了这个场景，但是testcase里有这个用例）
	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}
	if l < k {
		return head
	}

	// h，e用于记录已转换list部分的首和尾
	var h, e *ListNode
	for p, n := head, 0; p != nil; {
		n++
		p = p.Next
		if (p == nil) {
			break
		}
		if n%(k-1) == 0 {
			tmp := p.Next
			p.Next = nil
			sh, se := reverseList(head)
			se.Next = tmp
			// 如果h为空，说明是第一次转换；否则直接将转换后的list放到已转换列表的尾部就可以了
			if h == nil {
				h = sh
			} else {
				e.Next = sh
			}
			// 变更已转换list的尾部到se
			e = se
			head = tmp
			p = head
		}
	}
	return h

}



func main() {
	common.PrintList(reverseKGroup(common.Num2List(1, 2), 3))
	common.PrintList(reverseKGroup(common.Num2List(1, 2, 3), 3))
	common.PrintList(reverseKGroup(common.Num2List(1, 2, 3, 4), 3))
	common.PrintList(reverseKGroup(common.Num2List(1, 2, 3, 4, 5, 6), 3))
	common.PrintList(reverseKGroup(common.Num2List(1, 2, 3, 4, 5, 6, 7, 8), 3))
}
