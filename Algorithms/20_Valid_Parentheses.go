package main

import (
	"container/list"
	"fmt"
)

/*

20. Valid Parentheses
Easy

3141

149

Favorite

Share
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true

*/

//func isValid(s string) bool {
//	l := list.New()
//	for _, ch := range s {
//
//		if t := l.Back(); t == nil || t.Value != ch {
//			l.PushBack(ch)
//		} else {
//			l.Remove(t)
//		}
//	}
//	if l.Len() == 0 {
//		return true
//	} else {
//		return false
//	}
//}
//
//func main() {
//	fmt.Println(isValid("11"))
//	fmt.Println(isValid("12"))
//	fmt.Println(isValid("1221"))
//	fmt.Println(isValid("12211133"))
//}


func isValid(s string) bool {
	l := list.New()
	for _, ch := range s {
		if t := l.Back(); t != nil && ((t.Value == '(' && ch == ')') || (t.Value == '[' && ch == ']')|| (t.Value == '{' && ch == '}')){
			l.Remove(t)
		} else {
			l.PushBack(ch)
		}
	}

	if l.Len() == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid(")("))
	fmt.Println(isValid("){"))
}
