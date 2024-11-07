package problem_17

import "fmt"

func LetterCombinations_old(digits string) []string {

	var signalData = [][]string{
		{},
		{},
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}

	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return signalData[int(digits[0]-'0')]
	}

	var recalls []int // 记录要使用哪些层
	for _, item := range digits {
		recalls = append(recalls, int(item-'0'))
	}
	leftList := signalData[recalls[0]]
	rightList := signalData[recalls[1]]
	//for _, i := range recalls {
	for i := 1; i < len(recalls); i++ {
		rightList = signalData[recalls[i]]
		leftList = loop(leftList, rightList)
		fmt.Println("完成一轮")
	}

	return leftList
}

func loop(alphaList []string, betaList []string) []string {
	// 首先对alphaList做扩展，假设betaList 有n个，那就重复n次
	// 123 123 123 123
	// 444 555 666 777 重复k次，k是alphaList原本的长度
	k := len(alphaList)

	newBetaList := make([]string, len(betaList)*len(alphaList))
	alphaList = expand(alphaList, len(betaList))
	for j := range betaList {
		for t := 0; t < k; t++ {
			newBetaList[t+j*k] = betaList[j]
		}
	}

	for i := range alphaList {
		alphaList[i] = alphaList[i] + newBetaList[i]
	}
	return alphaList
}

func expand(alphaList []string, n int) []string {
	var result []string
	for i := 0; i < n; i++ {
		for _, item := range alphaList {
			result = append(result, item)
		}
	}
	return result
}
