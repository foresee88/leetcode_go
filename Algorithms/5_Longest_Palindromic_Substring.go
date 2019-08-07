//package ___Longest_Palindromic_Substring
package main

import "fmt"

/*

5. Longest Palindromic Substring
Medium

3878

370

Favorite

Share
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"

*/

func longestPalindrome(s string) string {
	l := len(s)
	subs := ""
	count := 0
	for i := 0; i < l; i++ {
		odd := getSame(s[0:i], s[i:l])
		single := getSame(s[0:i], s[i+1:l])

		if count < odd * 2 {
			count = odd *2
			subs = s[i-odd:i+odd]
		}

		if count < single * 2 + 1 {
			count = single * 2 + 1
			subs = s[i-single:i+1+single]
		}
	}
	return subs
}

func getSame(s0 string, s1 string) int{
	l := 0
	if len(s0) < len(s1) {
		l = len(s0)
	} else {
		l = len(s1)
	}

	i := len(s0)
	j := 0
	for ; j < l; j++ {
		if s0[i-1] != s1[j] {
			break
		}
		i--
	}
	return j
}

func main()  {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}