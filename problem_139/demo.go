package problem_139

/*
给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
*/
// 超出时间限制

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}
	// 整体
	x := false
	for _, word := range wordDict {
		if headCompare(s, word) {
			// 如果匹配成功了，那么就可以看继续匹配了
			x = x || backtracking(s[len(word):], wordDict)
		} else {
			// 如果匹配失败，那么结束
		}
	}
	return x
}

func backtracking(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	x := false
	for _, word := range wordDict {

		if headCompare(s, word) {
			// 如果匹配成功了，那么就可以看继续匹配了
			x = x || backtracking(s[len(word):], wordDict)
		} else {
			// 如果匹配失败，那么结束
		}

	}
	return x
}

func headCompare(s string, word string) bool {
	if len(s) < len(word) {
		return false
	}
	tmpS := s[0:len(word)]
	for i := range tmpS {
		if s[i] != word[i] {
			return false
		}
	}
	return true
}
