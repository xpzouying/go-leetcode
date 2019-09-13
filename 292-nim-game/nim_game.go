package nim_game

// --- 动态规划1 ---
// func canWinNim(n int) bool {
// 	if n < 4 {
// 		return true
// 	}

// 	// dp[i] 为数量为i时，我方的游戏结局
// 	// 记录 1 - n的记录表
// 	// 分配n+1空间。其中忽略0坐标，到n。
// 	dp := make([]bool, n+1)
// 	for i := 1; i < 4; i++ {
// 		dp[i] = true
// 	}

// 	for i := 4; i < n+1; i++ {

// 		// 我方分别拿1、2、3个石子后，对方的结果
// 		// 如果对方全胜，那么我方是失败；否则，我方胜利
// 		dp[i] = !(dp[i-1] && dp[i-2] && dp[i-3])
// 	}

// 	return dp[n]
// }

// --- 动态规划2: 优化空间 ---
// func canWinNim(n int) bool {
// 	if n < 4 {
// 		return true
// 	}

// 	// dp[i] 为数量为i时，我方的游戏结局
// 	// 记录 1 - n的记录表
// 	// 分配n+1空间。其中忽略0坐标，到n。
// 	dp := make([]bool, 3)
// 	for i := 0; i < 3; i++ {
// 		dp[i] = true
// 	}

// 	for i := 3; i < n; i++ {

// 		// 我方分别拿1、2、3个石子后，对方的结果
// 		// 如果对方全胜，那么我方是失败；否则，我方胜利
// 		current := !(dp[0] && dp[1] && dp[2])

// 		dp = append(dp[1:3], current)
// 	}

// 	return dp[2]
// }

func canWinNim(n int) bool {
	return (n % 4) != 0
}
