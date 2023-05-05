package algorithm

import (
	"strings"
	"fmt"
)

func checkNextJump(j int, currentChar string, pattern string) int{
	fmt.Println("This is j:", j)
	if strings.Contains(pattern, currentChar){
		leftString := pattern[:j]
		if(strings.Contains(leftString, currentChar)){
			lastIndex := strings.LastIndex(leftString, currentChar)
			fmt.Println("Case 1")
			return j - lastIndex
		}else{
			fmt.Println("Case 2")
			return 1;
		}
	}else{
		fmt.Println("Case 3")
		return len(pattern)
	}
}

func BoyerMoore(text, pattern string) int {
	n := len(text);
	m := len(pattern);

	if(m > n){
		return -1;
	}else{
		j := m - 1;
		i := m - 1;
		currentIndex := j
		test := 0
		for (currentIndex+1 <= n) {
			if(pattern[j] == text[i]){
				if(j == 0){
					return i;
				}else{
					i--;
					j--;
				}
			}else{
				nextJump := checkNextJump(j, string(text[i]), pattern)
				currentIndex = currentIndex + nextJump
				i = currentIndex
				j = m - 1;
				if test == 8{
					break
				}
				test++
			}
		}
		return -1;
	}
}