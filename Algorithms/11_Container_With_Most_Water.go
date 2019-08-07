//package _1__Container_With_Most_Water
package main

import (
	"fmt"
)

//Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.
//
//Note: You may not slant the container and n is at least 2.
//Example:
//
//Input: [1,8,6,2,5,4,8,3,7]
//Output: 49

func maxArea(height []int) int {
	max, area := 0, 0
	l := len(height)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if height[i] <= height[j] {
				area = height[i] * (j - i)
			} else {
				area = height[j] * (j - i)
			}
			if max < area {
				max = area
			}
		}
	}
	return max
}

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(arr))
}
