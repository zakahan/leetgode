package problem_139

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	m := wordBreak("bb", []string{"a", "b", "bbb"})
	fmt.Println(m)
}
