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
				// fmt.Println("This is current chat:", string(text[i]))
				// fmt.Println("This is the text that is being search before:", text[i-m+1:i])
				nextJump := checkNextJump(j, string(text[i]), pattern)
				fmt.Println("This is next jump value : ", nextJump)
				fmt.Println("This is next i marker before : ", i)
				currentIndex = currentIndex + nextJump
				i = currentIndex
				fmt.Println("This is next i marker after : ", i)
				// fmt.Println("This is the text that is being search after:", text[i-m+1:i+1])
				j = m - 1;
				fmt.Println("This is the length of text", n)
				if test == 8{
					break
				}
				test++
			}
		}
		return -1;
	}
}

// func main(){
// 	// text := "a pattern matching algorithm"
// 	text := "abacaabadcabacabaabb"
// 	// pattern := "rithm"
// 	pattern := "abacab"
// 	fmt.Println("Start")
// 	result := boyerMoore(text, pattern)
// 	fmt.Println("This is the result: ", result)
// }