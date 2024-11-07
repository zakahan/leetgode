package problem_3

func LengthOfLongestSubstring(s string) int {
	// 他是连续的
	if len(s) <= 1 {
		return len(s)
	}

	myMap := make(map[uint8]int)

	// 至少有两个
	var length = 0
	var i = 0           // 表示起点
	var j = 0           // 表示当前位置
	var flag = true     // 1表示可以滑动增加 0表示可以滑动减少
	var key uint8 = ' ' // 一开始应该是随便的
	for j < len(s) && i <= j {
		if flag {

			myMap[s[j]] += 1
			key = s[j]
			j++

		} else {
			myMap[s[i]] -= 1
			i++
		}

		// 检查一下是否需要改变flag
		if myMap[key] > 1 {
			flag = false
		} else {
			flag = true
			// 计算历史长度
			if j-i > length { // 只有在这个时候算的才是对的
				length = j - i
			}
		}

	}

	return length
}
