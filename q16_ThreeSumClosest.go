package leecode

import (
	"math"
	"sort"
)

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1,2,1,-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
*/

func ThreeSumClosest(nums []int, target int) int {
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	sort.Ints(nums)

	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		m := i + 1
		n := len(nums) - 1

		for {
			if m >= n {
				break
			}

			if nums[i]+nums[m]+nums[n] == target {
				return nums[i]+nums[m]+nums[n]
			}

			if nums[i]+nums[m]+nums[n] > target {
				n = n - 1
				for {
					if m >= n {
						if math.Abs(float64(nums[i]+nums[m]+nums[n+1]-target)) < math.Abs(float64(res-target)) {
							res = nums[i]+nums[m]+nums[n+1]
						}
						break
					}

					if nums[i]+nums[m]+nums[n] == target {
						return nums[i] + nums[m] + nums[n]
					}

					if nums[i]+nums[m]+nums[n] > target {
						n = n - 1
					} else {
						if math.Abs(float64(nums[i]+nums[m]+nums[n]-target)) < math.Abs(float64(res-target)) {
							res = nums[i]+nums[m]+nums[n]
						}
						if math.Abs(float64(nums[i]+nums[m]+nums[n+1]-target)) < math.Abs(float64(res-target)) {
							res = nums[i]+nums[m]+nums[n+1]
						}
						break
					}
				}
			} else {
				m = m + 1
				for {
					if m >= n {
						if math.Abs(float64(nums[i]+nums[m-1]+nums[n]-target)) < math.Abs(float64(res-target)) {
							res = nums[i]+nums[m-1]+nums[n]
						}
						break
					}

					if nums[i]+nums[m]+nums[n] == target {
						return nums[i] + nums[m] + nums[n]
					}

					if nums[i]+nums[m]+nums[n] < target {
						m = m + 1
					} else {
						if math.Abs(float64(nums[i]+nums[m]+nums[n]-target)) < math.Abs(float64(res-target)) {
							res = nums[i]+nums[m]+nums[n]
						}
						if math.Abs(float64(nums[i]+nums[m-1]+nums[n]-target)) < math.Abs(float64(res-target)) {
							res = nums[i]+nums[m-1]+nums[n]
						}
						break
					}
				}
			}
		}
	}

	return res
}
