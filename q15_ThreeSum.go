package leecode

import (
	"sort"
)

/*
q15

给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/

func ThreeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)

	var res [][]int
	length := len(nums)
	if nums[0] <= 0 && nums[length-1] >= 0 {
		for k, v := range nums {
			// in the case, [a,a,a...],
			// if a <= 0, only take care of the first left a
			// if a > 0, only take care of the last right a
			if k >= 1 && nums[k] <= 0 && nums[k] == nums[k-1] {
				continue
			}
			if k+1 < length && nums[k] > 0 && nums[k] == nums[k+1] {
				continue
			}

			i := k + 1
			j := length - 1
			for {
				if i >= j || v*nums[j] > 0 {
					break
				}

				sum := v + nums[i] + nums[j]
				if sum == 0 {
					res = append(res, []int{v, nums[i], nums[j]})
					i = i + 1
					j = j - 1
					for {
						if i < j &&  nums[j] == nums[j+1] {
							j = j - 1
						} else {
							break
						}
					}
					for {
						if i < j && nums[i] == nums[i-1] {
							i = i + 1
						} else {
							break
						}
					}
				} else if sum > 0 {
					j = j - 1
					for {
						if i < j && nums[j] == nums[j+1] {
							j = j - 1
						} else {
							break
						}
					}
				} else {
					i = i + 1
					for {
						if i < j && nums[i] == nums[i-1] {
							i = i + 1
						} else {
							break
						}
					}
				}
			}
		}
	}
	return res
}
