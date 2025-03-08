// -------------------------------------------------
// Package problem_438
// Author: hanzhi
// Date: 2025/2/5
// -------------------------------------------------

package problem_438

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)

	if len(p) > len(s) {
		return []int{} // 返回个空的，意思是没有
	}

	// 很明显的哈希表题目，算了人造哈希吧
	sMap := make([]int, 26)
	pMap := make([]int, 26) // pMap一次生成永久不变

	//for _, pSubString := range p {
	//	pMap[pSubString] += 1
	//}
	for i := 0; i < len(p); i++ {
		pSubString := p[i]
		pMap[int(pSubString-'a')] += 1

	}

	for i := 0; i < len(p)-1; i++ {
		sSubString := s[i]
		sMap[int(sSubString-'a')] += 1

	}

	// 打印
	//fmt.Println("当前p Map")
	//fmt.Println(pMap)
	//fmt.Println(sMap)

	// 遍历s
	for i := len(p) - 1; i < len(s); i++ {
		sMap[int(s[i]-'a')] += 1
		// 遍历两个map，对比
		var flag = 0
		for j := 0; j < 26; j++ {
			if sMap[j] != pMap[j] {
				// 证明失败了
				flag = 1
				break
			}
		}
		if flag == 0 {
			// 证明相等了，那么就 加进来
			//fmt.Printf("相等！当前subS:%v, p: %v\n", string(s[i-len(p)+1:i+1]), p)
			res = append(res, i-len(p)+1)
		}
		// 但是无论如何，下次开启遍历都会修改map
		sMap[int(s[i-len(p)+1]-'a')] -= 1

	}

	return res
}
