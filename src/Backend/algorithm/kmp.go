package algorithm

import (
	"strings"

	"github.com/agnivade/levenshtein"
)

func computeBorderFunction(pattern string) []int {
	n := len(pattern)
	borders := make([]int, n)
	for i, j := 1, 0; i < n; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = borders[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		borders[i] = j
	}
	return borders
}

// matchString returns true if the given pattern exists in the input string
func KMP(pattern, text string) int {
    pattern = strings.ToLower(pattern)
    text = strings.ToLower(text)
    // fmt.Println(computeBorderFunction(pattern))
	n, m := len(text), len(pattern)
	if m == 0 {
		return 0
	}
	borders := computeBorderFunction(pattern)
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && text[i] != pattern[j] {
			j = borders[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == m {
			return i-j+1
		}
	}
	return -1
}

func LevenshteinDistance (pattern string, text string) float64{
	pattern = strings.ToLower(pattern)
	text = strings.ToLower(text)
	distance := levenshtein.ComputeDistance(text, pattern)
	percentMatch := float64(len(text) - distance) / float64(len(text)) * 100
	return percentMatch
}

// func main(){
// 	Pattern := "Apa ibu kota indonsa"
// 	text := "Apa ibu kota Indonesia"
// 	// text := "abacaaabacaaxxxabacabxxx"
// 	// border := BorderFunction(Pattern)
// 	index := KMP(text, Pattern)
// 	fmt.Println(index)
// 	fmt.Println(levenshteinDistance(Pattern, text))
// }
