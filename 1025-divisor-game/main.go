package main

func divisorGame(n int) bool {
	switch n {
	case 1:
		return false
	case 2:
		return true
	}

	dp := make([]int, n+1)

	dp[2] = 1

	for i := 3; i <= n; i++ {

		for x := 1; x < i; x++ {
			// 异或操作，只需要找到一次，就能得到成功秘籍。
			// if dp[i] == 1 {
			// 	break
			// }

			// 爱丽丝选择的数字。
			// 如果满足整除条件，则x这个数时可以被选择的。
			// 那么，我们看看dp[i-x]是否为true（即1）
			if i%x == 0 {

				// 对方输的话，我们就能赢。
				if dp[i-x] == 0 {
					dp[i] = 1
					break
				}
			}
		}
	}

	return dp[n] == 1
}

/**

解题思路：

每次拿掉一个数 X，就相当于是出现了 (n-X)，
所以就看在DP[n-X]的时候，是否为true。我们只需要查找所有的DP[n-X]，X是选择可以被n整除的。

DP的初始化：

dp := make([]int, n+1)

// dp[1] = true

dp[2] = true

n == 3时，
爱丽丝：x只能选择 [1, 2]其中，但是有需要整除，所以只能选择1。n-->2
鲍勃：只能选择1。 n-->1
爱丽丝：输，false


**转移方程：**

dp[i] |= {
	dp[i-x]
}
如果dp[i] == true，则break；

*/
