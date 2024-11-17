package problem_76

// 本人手搓， leetcode判断是o(n^2),主要是map判断的环节时间复杂度比较大，然后我还要切割，也比较大，这两个地方可以砍应该
func MinWindow(s string, t string) string {
	//var record = []string{}
	var record = [][]int{}
	myMap := make(map[rune]int)
	for _, subT := range t {
		myMap[subT]++
	}

	localMap := make(map[rune]int)
	locationArr := make([]bool, len(s))
	var start = 0
	var end = 0

	for i, subS := range s {
		if myMap[subS] > 0 {
			locationArr[i] = true
		}
	}
	start = nextTrue(locationArr, -1)
	if start >= len(s) {
		return ""
	}
	localMap[rune(s[start])]++
	end = start
	for end <= len(s) {
		for mapIn(myMap, localMap) {
			record = append(record, []int{start, end + 1})
			localMap[rune(s[start])]--
			start = nextTrue(locationArr, start)
		}
		end = nextTrue(locationArr, end)
		if end == len(s) {
			break
		}
		localMap[rune(s[end])]++

	}
	var h = 0
	if len(record) == 0 {
		return ""
	}
	for i := range record {
		if record[i][1]-record[i][0] < record[h][1]-record[h][0] {
			h = i
		}
	}
	m := s[record[h][0]:record[h][1]]

	return m
}

func nextTrue(locationArr []bool, flag int) int {
	for flag = flag + 1; flag < len(locationArr) && locationArr[flag] == false; flag++ {

	}
	return flag
}

func mapIn(m1 map[rune]int, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		v2, exists := m2[k]
		if !exists || v1 > v2 { // v2是从父串儿里面选的，所以一定是v2要大于等于v1才对
			return false
		}
	}
	return true
}
