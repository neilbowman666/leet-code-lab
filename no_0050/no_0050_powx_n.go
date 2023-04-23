package no_0050

import "fmt"

// myPow
// ref: https://leetcode.cn/problems/powx-n/
// ref: https://blog.csdn.net/foolS22/article/details/125452877
func myPow(x float64, n int) float64 {
	fmt.Printf("myPow running, x:%v, n:%v\n", x, n)
	mount := x
	ans := 1.0

	// 正指数直接算出值，负指数算出来做分母即可
	N := n
	if n < 0 {
		N = -n
	}

	for {
		// golang 语法不支持，这里模拟 while N > 0
		if N <= 0 {
			break
		}

		ansUpdatedMsg := ""
		if N%2 == 1 {
			ans *= mount // 出现奇数指数则累积乘一次，最后一次迭代为1也会累积乘一次
			ansUpdatedMsg = ", ans updated"
		}

		N /= 2         // 每次迭代指数减半。最后一次的循环更新为0，将不进行下一次迭代。
		mount *= mount // 相乘迭代

		fmt.Printf("mount(powed):%v, N: %v, ans:%v%v\n", mount, N, ans, ansUpdatedMsg)
	}

	if n >= 0 {
		return ans
	} else {
		return 1.0 / ans
	}
}

func Test() {
	ansFunc := myPow
	fmt.Printf("leet code No.50 result sample 1: %v\n\n", ansFunc(2, 10))
	fmt.Printf("leet code No.50 result sample 2: %v\n\n", ansFunc(2.1, 3))
	fmt.Printf("leet code No.50 result sample 3: %v\n\n", ansFunc(2, -2))
}
