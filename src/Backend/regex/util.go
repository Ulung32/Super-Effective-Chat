package regex

import (
	"Backend/feature"
	"fmt"
	"log"
	"regexp"
)

func getDateQuery (query string) string {
	date := "\\b([0-2]?[0-9]|3[01])[./ -]([0]?[1-9]|1[012])[./ -]((19|20)\\d{2})\\b|\\b([0-2]?[0-9]|3[01]) (January|February|March|April|May|June|July|August|September|October|November|December) ((19|20)\\d{2})\\b"
	re := regexp.MustCompile(date)
	match := re.FindString(query)
	return match
}

func isDateQuery(query string) bool {
	date := "\\b([0-2]?[0-9]|3[01])[./ -]([0]?[1-9]|1[012])[./ -]((19|20)\\d{2})\\b|\\b([0-2]?[0-9]|3[01]) (January|February|March|April|May|June|July|August|September|October|November|December) ((19|20)\\d{2})\\b"
	re := regexp.MustCompile(date)
    match := re.MatchString(query)
	return match 
}

func getMathOperatorQuery(query string) string {
	mathOpr := "[([{\\s]*\\d+\\s*[)\\]}]*\\s*[+\\-*/]\\s*[([{\\s]*\\d+\\s*[)\\]}]*(\\s*[+\\-*/]\\s*[([{\\s]*\\d+\\s*[)\\]}]*)*"
	re := regexp.MustCompile(mathOpr)
	match := re.FindString(query)
	return match
}

func isMathOprQuery(query string) bool {
	mathOpr := "[([{\\s]*\\d+\\s*[)\\]}]*\\s*[+\\-*/]\\s*[([{\\s]*\\d+\\s*[)\\]}]*(\\s*[+\\-*/]\\s*[([{\\s]*\\d+\\s*[)\\]}]*)*"
	re := regexp.MustCompile(mathOpr)
    match := re.MatchString(query)
	return match
}
func isAddQuestionQuery(query string) bool {
	addQuestion := "tambah pertanyaan .* dengan jawaban .*"
	re := regexp.MustCompile(addQuestion)
    match := re.MatchString(query)
	return match
}
func isDeleteQuestionQuery(query string) bool {
	delQuestion := "hapus pertanyaan .*"
	re := regexp.MustCompile(delQuestion)
    match := re.MatchString(query)
	return match
}
func QueryClassification (query string) string {
	if(isDateQuery(query)){
		date := getDateQuery(query)
		fmt.Println(date)
		day, err := feature.GetDay(date)
		if(err != nil){
			log.Fatal(err)
		}
		answer := fmt.Sprintf("Date %s : %s", date, day)
		return answer
	}else if (isMathOprQuery(query)) {
		mathematicalExpression := getMathOperatorQuery(query)
		fmt.Println(mathematicalExpression)
		result, err := feature.MathematicalOperationSolver(mathematicalExpression)
		if(err != nil){
			log.Fatal(err)
		}
		answer := fmt.Sprintf("The result of the equation is %.3f", result)
		return answer
	}else if(isAddQuestionQuery(query)){
		return "4"
	}else if(isDeleteQuestionQuery(query)){
		return "5"
	}else{
		return "1"
	}
}
// func main(){
// 	query := "What is the day of 1 May 2023"
// 	fmt.Println(QueryClassification(query))
// }