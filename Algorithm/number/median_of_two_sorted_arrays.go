package main

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	midNums := make([]int, 0)
	i := 0
	j := 0
	length := len(nums1) + len(nums2)
	counter := 0
	isEVen := false
	if (length)/2 == 0 {
		isEVen = true
	}

	for i < len(nums1) || j < len(nums2) {
		if nums1[i] > nums2[j] {
			i++
			if isEVen && counter == length/2 {
				nu
			}
		} else {
			j++
		}
		counter++

	}
}