package problem_84

import (
	"fmt"
	"testing"
)

func TestLRLA(t *testing.T) {
	y := []int{5, 4, 1, 2}
	x := largestRectangleArea(y)
	fmt.Println(x)
}
