package problem_48

/*
123         741
456  -》     852
789         963

00 -> nn -> 0n

01 -> 12
*/

func rotate(matrix [][]int) {
	// 好像可以两次处理，先沿副对角反转 然后再上下对调
	n := len(matrix) - 1
	for i := 0; i < len(matrix); i++ {
		// x[i,j] -> x[n-i, n-j]
		for j := 0; j < len(matrix)-i; j++ {
			matrix[i][j], matrix[n-j][n-i] = matrix[n-j][n-i], matrix[i][j]
			//matrix[i][j] = -1
		}
	}

	//for i := range matrix {
	//	for j := range matrix {
	//		fmt.Print(" ", matrix[i][j])
	//	}
	//	fmt.Println()
	//}

	for i := 0; i < len(matrix)/2; i++ {
		// x[i,j] -> x[n-i, n-j]
		for j := 0; j < len(matrix); j++ {
			matrix[i][j], matrix[n-i][j] = matrix[n-i][j], matrix[i][j]
		}
	}

	//for i := range matrix {
	//	for j := range matrix {
	//		fmt.Print(" ", matrix[i][j])
	//	}
	//	fmt.Println()
	//}
}
