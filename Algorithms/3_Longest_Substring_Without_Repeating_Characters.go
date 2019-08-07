//package ___Longest_Substring_Without_Repeating_Characters
package main

import (
	"strings"
	"fmt"
)

/*

3. Longest Substring Without Repeating Characters
Medium

5824

333

Favorite

Share
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	subs := ""
	count := 0
	for _, ch := range s {
		if !strings.Contains(subs, string(ch)) {
			subs += string(ch)
		} else {
			subs = strings.Split(subs,string(ch))[1]
			subs += string(ch)
		}
		if count < len(subs) {
			count = len(subs)
		}
	}
	return count
}

func main(){
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
