package problem_48

func Rotate(matrix [][]int) {
	// 分四块旋转法
	// 只需要对矩阵的左上角做遍历就好了
	// [i,j] -> [j,n-i] -> [n-i,n-j] -> [n-j,i] 交换，然后n-列 wow
	// 其实就是法1的循环
	// 0,1  -> 1,3  -> 3,2 -> 2,0
	n := len(matrix) - 1
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < (len(matrix)+1)/2; j++ {
			matrix[i][j], matrix[j][n-i], matrix[n-i][n-j], matrix[n-j][i] =
				matrix[n-j][i], matrix[i][j], matrix[j][n-i], matrix[n-i][n-j]
		}
	}
	//
	//for i := range matrix {
	//	for j := range matrix {
	//		fmt.Print(" ", matrix[i][j])
	//	}
	//	fmt.Println("|")
	//}
}
