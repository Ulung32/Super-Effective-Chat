package main

import (
	"fmt"
	"strings"

	"github.com/agnivade/levenshtein"
)

func BorderFunction(Pattern string)[] int{
	var n int
	n = len(Pattern)
	border := make([]int, n)
    for i, j := 1, 0; i < n; i++ {
        for j > 0 && Pattern[i] != Pattern[j] {
            j = border[j-1]
        }
        if Pattern[i] == Pattern[j] {
            j++
        }
        border[i] = j
    }

    return border
}

func KMP(pattern string, text string) int{
	pattern = strings.ToLower(pattern)
	text = strings.ToLower(text)
	miss := false
	j := 0
	border := BorderFunction(pattern)
	// fmt.Println(border)

	for i := 0; i < len(text); i++{
		if j == len(pattern){
			return i - j 
		}else{
			if text[i] == pattern[j]{
				miss = false
			}else{
				miss = true
			}
		}
		if miss == true {
			if j!=0{
				j = border[j-1]
			}
			
		}else {
			j = j + 1
		}
	}
	if miss == true {
		return -1
	}else{
		return len(text) - len(pattern)
	}
}

func levenshteinDistance (pattern string, text string) float64{
	pattern = strings.ToLower(pattern)
	text = strings.ToLower(text)
	distance := levenshtein.ComputeDistance(text, pattern)
	percentMatch := float64(len(text) - distance) / float64(len(text)) * 100
	return percentMatch
}

func main(){
	Pattern := "Apa ibu kota indonsa"
	text := "Apa ibu kota Indonesia"
	// text := "abacaaabacaaxxxabacabxxx"
	// border := BorderFunction(Pattern)
	index := KMP(text, Pattern)
	fmt.Println(index)
	fmt.Println(levenshteinDistance(Pattern, text))
}
