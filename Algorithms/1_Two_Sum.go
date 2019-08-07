//package main
package main

import (
	"fmt"
)

/*

1. Two Sum
Easy

11137

381

Favorite

Share
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return []int{i,j}
			}
		}
	}
	return []int{0,0}
}

func main(){
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
}
