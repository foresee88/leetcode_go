//package main
package main

import (
	"github.com/foresee88/leetcode_go/Algorithms/common"
)

/*

2. Add Two Numbers
Medium

5468

1405

Favorite

Share
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = common.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{Val: (l1.Val + l2.Val) % 10, Next: nil}
	up := (l1.Val + l2.Val) / 10
	l := ret
	l1 = l1.Next
	l2 = l2.Next

	for ; l1 != nil || l2 != nil || up != 0; {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		num := val1 + val2 + up
		up = num / 10
		l.Next = &ListNode{Val: num % 10, Next: nil}
		l = l.Next
	}

	return ret
}


func main() {
	common.PrintList(addTwoNumbers(common.Num2List(2, 4, 3), common.Num2List(5, 6, 4)))
	common.PrintList(addTwoNumbers(common.Num2List(7), common.Num2List(5, 0, 5)))
}
