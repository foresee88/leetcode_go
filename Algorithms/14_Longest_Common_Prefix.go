package main

import "fmt"

/*

14. Longest Common Prefix
Easy

1452

1381

Favorite

Share
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	l := len(strs[0])
	for i := 1; i< len(strs); i++ {
		if l > len(strs[i]) {
			l = len(strs[i])
		}
	}
	if l == 0 {
		return ""
	}

	i := 0
	for ; i < l; i++ {
		base := strs[0][i]
		j := 1
		for ; j < len(strs); j++ {
			if base != strs[j][i] {
				break
			}
		}
		if j != len(strs) {
			break
		}
	}

	return strs[0][0:i]

}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{}))
	fmt.Println(longestCommonPrefix([]string{"", "ab"}))
}
