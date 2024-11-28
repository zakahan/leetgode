package problem_139

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	m := wordBreak("leetcode", []string{"leet", "code"})
	fmt.Println(m)
}
