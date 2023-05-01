package controller

import (
	"Backend/algorithm"
	"Backend/models"
	"fmt"
)
type QueryProcessor struct {
	QnAList []models.QnA
	// Query string
}

var Processor QueryProcessor

func (p QueryProcessor) QuerySearch(method, query string) (int, float64) {
	// if(method == "kmp"){

	// }
	var index int
	var similarity float64

	similarity = 0
	index = -1
	for i:=0; i<len(p.QnAList); i++{
		fmt.Println(query, p.QnAList[i].Question)
		var result int;
		if method == "kmp"{
			result = algorithm.KMP(query, p.QnAList[i].Question)
		}else{
			result = algorithm.BoyerMoore(query, p.QnAList[i].Question)
		}
		if result != -1{
			index = i
			similarity = 100
			return index, similarity
		}else{
			temp := algorithm.LevenshteinDistance(query, p.QnAList[i].Question)
			if(temp > similarity) {
				similarity = temp
				index = i
			}
		}
		// fmt.Println(index, similarity)
	}
	return index, similarity
}