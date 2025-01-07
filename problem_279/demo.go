// -------------------------------------------------
// Package problem_279
// Author: hanzhi
// Date: 2025/1/7
// -------------------------------------------------

package problem_279

import "fmt"

func getSq() {
	i := 1
	var z []int = []int{}

	for i*i < 1000 {
		z = append(z, i*i)
		i = i + 1
	}
	fmt.Print("{")
	for i := range z {
		fmt.Printf(`%v,`, z[i])
	}
	fmt.Println("}")
}

func numSquares(n int) int {
	list := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100, 121, 144, 169, 196, 225, 256, 289, 324, 361, 400, 441, 484, 529, 576, 625, 676, 729, 784, 841, 900, 961}
	fmt.Println(list)
	return list[0]
}
