package main

import (
	"fmt"
	"sort"
)

/*

18. 4Sum
Medium

1163

234

Favorite

Share
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:

Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
[-1,  0, 0, 1],
[-2, -1, 1, 2],
[-2,  0, 0, 2]
]
*/

func same(s, d []int) bool {
	if len(s) != len(d) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != d[i] {
			return false
		}
	}
	return true
}

func fourSum(nums []int, target int) [][]int {
	var ret, arr [][]int
	sort.Ints(nums)
	l := len(nums)
	for i := 0; i < l-3; i++ {
		for j := i + 1; j < l-2; j++ {
			for k := j + 1; k < l-1; k++ {
				for m := k + 1; m < l; m++ {
					if nums[i]+nums[j]+nums[k]+nums[m] == target {
						arr = append(arr, []int{nums[i], nums[j], nums[k], nums[m]})
					}
				}
			}
		}
	}

	//去重
	la := len(arr)
	for i := 0; i < la; i++ {
		j := i + 1
		for ; j < la; j++ {
			if same(arr[i], arr[j]) {
				break
			}
		}
		if j >= la {
			ret = append(ret, arr[i])
		}
	}

	return ret
}

func main() {
	aa := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Print(aa)
	aa1 := fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0)
	fmt.Print(aa1)
}
