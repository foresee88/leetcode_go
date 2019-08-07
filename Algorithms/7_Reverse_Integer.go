//package ___Reverse_Integer
package main

import (
	"fmt"
	"math"
)

/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321

Example 2:

Input: -123
Output: -321

Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

func reverse(x int) int {
	if x < -0x7fffffff || x > (0x7fffffff-1) {
		return 0
	}
	minus := false
	if x < 0 {
		minus = true
		x = int(math.Abs(float64(x)))
	}
	ret := 0
	for x > 0 {
		ret = ret*10 + (x % 10)
		x = x / 10
	}
	if minus {
		ret = 0 - ret
	}
	if ret < -0x7fffffff || ret > (0x7fffffff-1) {
		return 0
	}
	return ret
}

func main() {
	t1, t2, t3, t4 := 123, -123, 120, 1534236469
	fmt.Println(reverse(t1))
	fmt.Println(reverse(t2))
	fmt.Println(reverse(t3))
	fmt.Println(reverse(t4))

}
