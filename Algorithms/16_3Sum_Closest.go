package main

import (
	"sort"
	"math"
	"fmt"
)

/*
16. 3Sum Closest
Medium

1197

87

Favorite

Share
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
*/


func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	num := nums[0]+nums[1]+nums[2]
	offset := math.Abs((float64)(num-target))

	i,j,k,l := 0,0,0,len(nums)
	for ;i < l-2;i++{
		for j=i+1;j<l-1;j++{
			for k=j+1;k<l;k++ {
				t := math.Abs((float64)((nums[i]+nums[j]+nums[k])-target))
				if offset > t {
					offset = t
					num = nums[i]+nums[j]+nums[k]
				}
			}
		}
	}
	return num

}

func main(){
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}