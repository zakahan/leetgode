package problem_139

/*
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
*/
// 破防了，这测试用例实在是太全面了，用递归记忆化搜索方法，会超时，先留着把，不知道怎么处理了

func wordBreakF(s string, wordDict []string) bool {
	// 整一个字典值遍历，如果不存在就剔除（小聪明被ban了）
	//letterDict := make(map[uint8]bool)
	//for i := range wordDict {
	//	for j := range wordDict[i] {
	//		letterDict[wordDict[i][j]] = true
	//	}
	//}
	//
	//for i := range s {
	//	if !letterDict[s[i]] {
	//		return false
	//	}
	//}

	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}
	return backtracking(s, wordMap)
}

func backtracking(s string, wordMap map[string]bool) bool {
	if len(s) == 0 {
		return true
	}
	x := false
	for i := 0; i <= len(s); i++ {
		if headCompare(s[0:i], wordMap) {
			// 如果匹配成功了，说明之前的存在，那就只要匹配后面的就可以了。
			x = x || backtracking(s[i:], wordMap)
			//if x{
			//	return x
			//}
		}
	}
	return x
}

func headCompare(s string, wordMap map[string]bool) bool {
	//fmt.Println(s)
	return wordMap[s]
}
