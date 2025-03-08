package problem_96

/*
一个的话，一种
两个的话 两种（1头2右  2头1左
三个的话 在2的基础上，2*2 + 1 = 5
四个的话 在三的基础上  1个头节点， 左0右3 左1右2 左2右1 左3右0
1*5 + 1*2 +  2*1 + 5*1 =
*/
func numTreesBT(n int) int {
	return backtracking(n)
}

func backtracking(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		// 常规数值
		x := 0
		for i := 0; i <= n-1; i++ {
			//fmt.Println("G[", i, "]*", "G[", n-1-i, "]")
			x += backtracking(i) * backtracking(n-1-i)
		}
		return x
	}

}
