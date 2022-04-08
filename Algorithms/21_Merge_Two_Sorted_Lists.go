/*
21. Merge Two Sorted Lists
Easy

11651

1059

Add to List

Share
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.



Example 1:


Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]


Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/

package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var l, c *ListNode
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		l, c = list2, list2
		list2 = list2.Next
	} else {
		l, c = list1, list1
		list1 = list1.Next
	}
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			c.Next = list2
			c = c.Next
			list2 = list2.Next
		} else {
			c.Next = list1
			c = c.Next
			list1 = list1.Next
		}
	}
	if list1 != nil{
		c.Next = list1
	} else if list2 != nil{
		c.Next = list2
	}

	return l

}

func main() {
	PrintList(mergeTwoLists(Num2List(1, 2, 4), Num2List(1, 3, 4)))
	PrintList(mergeTwoLists(Num2List(2), Num2List(1)))
	PrintList(mergeTwoLists(nil, nil))
}
