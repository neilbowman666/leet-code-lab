package no_0015

import "leet-code-lab/utils"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	utils.BubbleSort(&nums, true)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for {
			if left >= right {
				break
			}
			if nums[i]+nums[left]+nums[right] < 0 { // 需要更大的数，那就要往后移动 left
				// 去重，移动过后值与之前一样那就没有意义
				// 下面的循环表示：继续移动直到移动后的 left 值与之前的 left 值不相同
				for {
					left++
					if !(left < right && nums[left] == nums[left-1]) { // 取 left-1 是安全的
						break
					}
				}
			} else if nums[i]+nums[left]+nums[right] > 0 { // 需要更小的数，那就要往前移动 right
				// 去重，移动过后值与之前一样那就没有意义
				// 下面的循环表示：继续移动直到移动后的 right 值与之前的 right 值不相同
				for {
					right--
					if !(left < right && nums[right] == nums[right+1]) { // 取 right+1 是安全的
						break
					}
				}
			} else { // 已经找到想要的数，记录到结果集，并同时向后移动 left 和向前移动 right
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// 去重
				for {
					if !(left < right && nums[right] == nums[right-1]) {
						break
					}
					right--
				}

				// 去重
				for {
					if !(left < right && nums[left] == nums[left+1]) {
						break
					}
					left++
				}

				right--
				left++
			}
		}
	}
	return result
}
