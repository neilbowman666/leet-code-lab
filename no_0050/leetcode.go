package no_0050

import "fmt"

func quickM(x float64, N int) float64 {
	if N == 0 { // 退出递归的条件
		return 1.0
	}
	y := quickM(x, N/2)
	if N%2 == 0 {
		return y * y
	} else {
		return y * y * x
	}
}

// myPow1
// ref: https://leetcode.cn/problems/powx-n/
// ref: https://blog.csdn.net/foolS22/article/details/125452877
func myPow1(x float64, n int) float64 {
	if n > 0 {
		return quickM(x, n)
	} else {
		return 1.0 / quickM(x, n)
	}
}

// myPow2
// 递归改迭代方法：https://blog.csdn.net/u014698348/article/details/53216934
func myPow2(x float64, n int) float64 {
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
		// 每次循环 N = N /2，最后一次循环时 N = 1，更新为 0 后将不再下一次循环。
		if N <= 0 {
			fmt.Printf("now N is: %v, stopped.\n", N)
			break
		}

		ansUpdatedMsg := ""
		oldAns := ans
		if N%2 == 1 {
			ans *= mount // 出现奇数指数则累积乘一次，最后一次迭代为1也会累积乘一次
			ansUpdatedMsg = ", ans updated"
		}

		oldN := N
		oldMount := mount
		mount *= mount // 相乘迭代
		N /= 2         // 每次迭代指数减半。最后一次的循环更新为0，将不进行下一次迭代。

		fmt.Printf("oldAns:%v,\t oldMount:%v,\t mount(powed):%v,\t oldN:%v,\t N(halfed):%v,\t ans:%v%v.\n", oldAns, oldMount, mount, oldN, N, ans, ansUpdatedMsg)
	}

	if n >= 0 {
		return ans
	} else {
		return 1.0 / ans
	}
}
