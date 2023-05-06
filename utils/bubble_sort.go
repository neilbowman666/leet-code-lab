package utils

// BubbleSort 冒泡排序
// 从前往后遍历整个数组，最大值或最小值排最前面
// 遍历数组是第一层循环，找最大值是第二层循环
func BubbleSort(originalIntegers *[]int, asc bool) {
	for i := 0; i < len(*originalIntegers)-1; i++ {
		for j := i + 1; j < len(*originalIntegers); j++ {
			if asc { // 升序
				if (*originalIntegers)[i] > (*originalIntegers)[j] {
					temp := (*originalIntegers)[i]
					(*originalIntegers)[i] = (*originalIntegers)[j]
					(*originalIntegers)[j] = temp
				}
			} else { // 降序
				if (*originalIntegers)[i] <= (*originalIntegers)[j] {
					temp := (*originalIntegers)[i]
					(*originalIntegers)[i] = (*originalIntegers)[j]
					(*originalIntegers)[j] = temp
				}
			}
		}
	}
}
