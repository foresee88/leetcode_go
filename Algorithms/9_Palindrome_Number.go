//package ___Palindrome_Number
package main

import (
	"fmt"
	"strconv"
)

/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
Follow up:

Coud you solve it without converting the integer to a string?


*/

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	i, l := 0, len(str)
	for ; i <= l/2; i++ {
		if str[i] != str[l-i-1] {
			return false
		}
	}
	return true
}

func main() {
	t1, t2, t3 := 121, -121, 10
	fmt.Println(isPalindrome(t1))
	fmt.Println(isPalindrome(t2))
	fmt.Println(isPalindrome(t3))
}
