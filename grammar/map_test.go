package grammar

import (
	"fmt"
	"testing"
)

func mapToSlice(m map[int]string) []string {
	s := make([]string, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}
func TestMapSlice(t *testing.T) {
	var m = map[int]string{1: "hello", 2: "word"}
	fmt.Print(mapToSlice(m))
	fmt.Print([]string{"1", "2"})
}
