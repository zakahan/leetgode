package problem_72

/*
如何理解dp
别几把听什么对word1和对word2的等价了，两个变动来理解起来太傻逼了。
dp表想象有一个游标
游标最开始在开头位置（空字符）
都是对word1做操作！
比如abc  de
往右是在这个位置添加word2对应的字符   dp[0][1] 就是将abc变成  dabc
往下是在这个位置删除word1对应的字符   dp[1][0] 就是将abc变成 bc
往右下是替换word1的[i]为word2的[j]  dp[1][1] 就是删掉a， 然后添加d  变成了 dbc， 关键在于状态转移


如何理解状态转移方程
知道定义了就好理解了，不多说了，反正就是dp[i,j] 能有三种来源无非就是 删除i号，增加j号，以及先删后改（这仨都是一次操作，只不过为了理解）
*/

func minDistance(word1 string, word2 string) int {

	m := len(word1) + 1
	n := len(word2) + 1

	if len(word1) == 0 || len(word2) == 0 {
		return len(word1) + len(word2)
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 0

	for i := range dp {
		dp[i][0] = i
	}
	for j := range dp[0] {
		dp[0][j] = j
	}

	// 以上完成初始化了
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i-1] == word2[j-1] { // 为什么要-1,因为要适配开始游标在最开头，那个位置理论来说是-1没办法，挪一个位置
				//dp[i][j] = minInt3(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
				dp[i][j] = dp[i-1][j-1] // 如果相同，不用管，大概就是这个概念
			} else {
				dp[i][j] = minInt3(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m-1][n-1]
}

func minInt3(a int, b int, c int) int {
	if a < b && a < c {
		return a
	} else if b < a && b < c {
		return b
	} else {
		return c
	}
}
