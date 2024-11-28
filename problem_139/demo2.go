package problem_139

// 这种题有一类共同点，就是只需要记住，我能到达某个地方就行了，后面的不用管前面是怎么到的！
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}
	dp := make([]bool, len(s)+1)

	dp[0] = true
	for i := 0; i <= len(s); i++ {
		// 来中途分割点
		for j := 0; j < i; j++ {
			dp[i] = dp[i] || (dp[j] && isHere(s[j:i], wordMap))
			if dp[i] {
				break
			}
		}
	}
	return dp[len(s)]
}

func isHere(s string, wordMap map[string]bool) bool {
	//fmt.Println(s)
	return wordMap[s]
}
