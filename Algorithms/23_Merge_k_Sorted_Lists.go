package main

import (
	"Algorithms/common"
)

/*

23. Merge k Sorted Lists
Hard

2641

171

Favorite

Share
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
1->4->5,
1->3->4,
2->6
]
Output: 1->1->2->3->4->4->5->6
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = common.ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	// 预处理，清除空的列表
	ls := make([]*ListNode, 0)
	for _, l := range lists {
		if l != nil {
			ls = append(ls, l)
		}
	}

	if len(ls) == 0 {
		return nil
	}

	ret := ls[0]
	for i := 1; i < len(ls); i++ {
		d := ret
		s := ls[i]
		for ; s != nil; {
			if s.Val < d.Val { // s小于d，插到d的前方，需要变更ret
				t := s.Next
				s.Next = d
				d = s
				ret = d
				s = t
			} else if s.Val >= d.Val && d.Next != nil && s.Val < d.Next.Val { // s大于等于d并小于d.Next,插d的后面，d往后移动到d.Next
				t := s.Next
				s.Next = d.Next
				d.Next = s
				s = t
				d = d.Next
			} else if s.Val >= d.Val && d.Next == nil { // s小于d，并且d是列表最后一个，s直接拼接到d后面即可
				d.Next = s
				break
			} else { // s大于d，但不小于d.Next，且d.Next不等于nil，d往后移动到d.Next
				d = d.Next
			}
		}
	}
	return ret

}


func main() {
	lists := []*ListNode{common.Num2List(1, 4, 5), common.Num2List(1, 3, 4), common.Num2List(2, 6)}
	common.PrintList(mergeKLists(lists))

	lists1 := []*ListNode{common.Num2List(1, 4, 5), common.Num2List(2, 3)}
	common.PrintList(mergeKLists(lists1))

	lists2 := []*ListNode{nil, common.Num2List(1)}
	common.PrintList(mergeKLists(lists2))

	lists3 := []*ListNode{}
	common.PrintList(mergeKLists(lists3))

	lists4 := []*ListNode{common.Num2List(1), common.Num2List(0)}
	common.PrintList(mergeKLists(lists4))

	lists5 := []*ListNode{common.Num2List(-1, -1, -1), common.Num2List(-2, -2, -1)}
	common.PrintList(mergeKLists(lists5))

}
