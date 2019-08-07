package common

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 将入参数字转换成ListNode列表
func Num2List(n ...int) *ListNode {
	ret := ListNode{Val: n[0], Next: nil}
	p := &ret
	for i := 1; i < len(n); i++ {
		p.Next = &ListNode{Val: n[i], Next: nil}
		p = p.Next
	}
	return &ret
}


// 打印ListNode列表，输出格式"1 -> 2 -> 3"
func PrintList(l *ListNode) {
	var s string
	for ; l != nil; l = l.Next {
		if s == "" {
			s = fmt.Sprint(l.Val)
		} else {
			s = fmt.Sprint(s, " -> ", l.Val)
		}
	}
	fmt.Println(s)
}


