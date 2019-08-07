//package _0__Regular_Expression_Matching
package main

import (
	"fmt"
	"strings"
)

/*
Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like . or *.
Example 1:

Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the precedeng element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
Example 4:

Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore it matches "aab".
Example 5:

Input:
s = "mississippi"
p = "mis*is*p*."
Output: false
*/


//todo
func isMatch(s string, p string) bool {

	ss := getRegs(p)

	idx := 0
	for i := 0; i < len(s); i++ {
		if idx >= len(ss) {
			return false
		}
		if len(ss[idx]) > 1 {
			if ss[idx][0] == '.' || (s[i] == ss[idx][0]) {
				continue
			} else {
				idx++
				i--
				continue
			}
		} else if ss[idx][0] == '.' || (s[i] == ss[idx][0]) {
			idx++
			continue
		} else {
			return false
		}

	}

	if idx < len(ss) {
		return false
	}

	return true
}

// 分解出可以用于匹配的子字符串
func getRegs(p string) []string {

	// 空字符处理
	if len(p) == 0 {
		return []string{}
	}

	// 去掉头部"*"
	p = strings.TrimPrefix(p, "*")

	// 去掉字符串中的连续重复的"*"
	s := string(p[0])
	for i := range p {
		if i > 0 {
			if p[i] == '*' && p[i] == p[i-1] {
				continue
			}
			s += string(p[i])
		}
	}

	//// 按"*"分解出可以用于匹配的子字符串
	//var strs []string
	//ss := strings.SplitAfter(s, "*")
	//for _, s := range ss {
	//	if len(s) > 2 && strings.HasSuffix(s, "*") {
	//		strs = append(strs, s[:len(s)-2])
	//		strs = append(strs, s[len(s)-2:])
	//	} else {
	//		strs = append(strs, s)
	//	}
	//}

	var strs []string
	j := 0
	for ; j < (len(s) - 1); j++ {
		if s[j+1] == '*' {
			strs = append(strs, s[j:j+2])
			j++
		} else {
			strs = append(strs, string(s[j]))
		}
	}
	if j < len(s) {
		strs = append(strs, string(s[j]))
	}
	fmt.Println(strs)

	return strs
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))

	fmt.Println(isMatch("ab", ".*c"))
}
