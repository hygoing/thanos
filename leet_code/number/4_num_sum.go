package main

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
 */

import "sort"

func fourSum(nums []int, target int) [][]int {
	answers := make([][]int, 0)
	if len(nums) < 4 {
		return answers
	}

	sort.Ints(nums)
	for i := 0; i <= len(nums)-4; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j <= len(nums)-3; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			L := j + 1
			R := len(nums) - 1
			for L < R {
				sum := nums[i] + nums[j] + nums[L] + nums[R]
				if sum == target {
					answers = append(answers, []int{nums[i], nums[j], nums[L], nums[R]})
					for L < R && nums[L] == nums[L+1] {
						L++
					}
					for R > L && nums[R] == nums[R-1] {
						R--
					}
					L++
					R--
				}
				if sum > target {
					R --
				}
				if sum < target {
					L++
				}

			}
		}
	}
	return answers
}
