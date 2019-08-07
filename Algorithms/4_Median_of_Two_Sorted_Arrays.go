//package ___Median_of_Two_Sorted_Arrays

package main

import (
	"sort"
	"fmt"
)

/*

4. Median of Two Sorted Arrays
Hard

4581

640

Favorite

Share
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5

*/


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	m := len(nums)
	if m%2 == 1 {
		return float64(nums[m/2])
	} else {
		return (float64(nums[m/2-1]) + float64(nums[m/2]))/2
	}
}

func main()  {
	fmt.Println(findMedianSortedArrays([]int{1,3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1,2}, []int{3,4}))
}