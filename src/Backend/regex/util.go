package regex

import (
	"Backend/feature"
	"fmt"

	// "log"
	"regexp"
	"strings"
)

func getDateQuery (query string) string {
	date := "\\b([0-2]?[0-9]|3[01])[./ -]([0]?[1-9]|1[012])[./ -]((19|20)\\d{2})\\b|\\b([0-2]?[0-9]|3[01]) (January|February|March|April|May|June|July|August|September|October|November|December) ((19|20)\\d{2})\\b"
	re := regexp.MustCompile(date)
	match := re.FindString(query)
	return match
}

func isDateQuery(query string) (bool, bool) {
	dateFixedRegex := "(day|date|Day|Date)"
	date := "\\b([0-2]?[0-9]|3[01])[./ -]([0]?[1-9]|1[012])[./ -]((19|20)\\d{2})\\b|\\b([0-2]?[0-9]|3[01]) (January|February|March|April|May|June|July|August|September|October|November|December) ((19|20)\\d{2})\\b"
	re := regexp.MustCompile(date)
	reFixed := regexp.MustCompile(dateFixedRegex)
    match := re.MatchString(query)
	fixedMatch := reFixed.MatchString(query)
	return match, fixedMatch
}

func getMathOperatorQuery(query string) string {
	mathOpr := "[([{\\s]*\\d+\\s*[)\\]}]*\\s*[+\\-*/]\\s*[([{\\s]*\\d+\\s*[)\\]}]*(\\s*[+\\-*/]\\s*[([{\\s]*\\d+\\s*[)\\]}]*)*"
	re := regexp.MustCompile(mathOpr)
	match := re.FindString(query)
	return match
}

func isMathOprQuery(query string) (bool, bool) {
	equationFixedRegex := "(equation|Equation)"
	mathOpr := "[([{\\s]*\\d+\\s*[)\\]}]*\\s*[+\\-*/]\\s*[([{\\s]*\\d+\\s*[)\\]}]*(\\s*[+\\-*/]\\s*[([{\\s]*\\d+\\s*[)\\]}]*)*"
	re := regexp.MustCompile(mathOpr)
	reFixed := regexp.MustCompile(equationFixedRegex)
    match := re.MatchString(query)
	fixedMatch := reFixed.MatchString(query)
	return match, fixedMatch
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

func SplitQuestion(query string) []string {
	sep := []string{"?", "!", "\n"}
	question := strings.FieldsFunc(query, func(r rune) bool {
		for _, s := range sep {
			if r == []rune(s)[0] {
				return true
			}
		}
		return false
	})
	return question
}
func QueryClassification (query string) string {
	// var answer string;
	// for _, question := range questions{
	// 	if(question != ""){
	// 		fmt.Println("This is the question", question)
	// 		tempAnswer := ""
	// 		dateMatch, commandDateMatch := isDateQuery(question)
	// 		equationMatch, commandEquationMatch := isMathOprQuery(question)
	// 		if((commandDateMatch || dateMatch) && !commandEquationMatch){
	// 			if(dateMatch){
	// 				date := getDateQuery(question)
	// 				fmt.Println("This is date", date)
	// 				day, err := feature.GetDay(date)
	// 				if(err != nil){
	// 					tempAnswer = "Invalid Date\n"
	// 				}else{
	// 					tempAnswer = fmt.Sprintf("Date %s : %s\n", date, day)
	// 				}
	// 			}else{
	// 				tempAnswer = "Invalid Date\n"
	// 			}
	// 		}else if (commandEquationMatch || equationMatch) {
	// 			if(equationMatch){
	// 				mathematicalExpression := getMathOperatorQuery(question)
	// 				fmt.Println("This is mathematical expression", mathematicalExpression)
	// 				result, err := feature.MathematicalOperationSolver(mathematicalExpression)
	// 				if(err != nil){
	// 					tempAnswer = "Invalid Syntax\n"
	// 				}else{
	// 					tempAnswer = fmt.Sprintf("The result of the equation is %.3f\n", result)
	// 				}
	// 			}else{
	// 				tempAnswer = "Invalid Syntax\n"
	// 			}
	// 		}else if(isAddQuestionQuery(question)){
	// 			return "4"
	// 		}else if(isDeleteQuestionQuery(question)){
	// 			return "5"
	// 		}else{
	// 			return "1"
	// 		}
	// 		answer = answer + tempAnswer
	// 	}
	// }
	// return answer
	dateMatch, commandDateMatch := isDateQuery(query)
	equationMatch, commandEquationMatch := isMathOprQuery(query)
	if((commandDateMatch || dateMatch) && !commandEquationMatch){
		if(dateMatch){
			date := getDateQuery(query)
			fmt.Println("This is date", date)
			day, err := feature.GetDay(date)
			if(err != nil){
				return "Invalid Date\n"
			}else{
				return fmt.Sprintf("Date %s : %s\n", date, day)
			}
		}else{
			return "Invalid Date\n"
		}
	}else if (commandEquationMatch || equationMatch) {
		if(equationMatch){
			mathematicalExpression := getMathOperatorQuery(query)
			fmt.Println("This is mathematical expression", mathematicalExpression)
			result, err := feature.MathematicalOperationSolver(mathematicalExpression)
			if(err != nil){
				return "Invalid Syntax\n"
			}else{
				return fmt.Sprintf("The result of the equation is %.3f\n", result)
			}
		}else{
			return "Invalid Syntax\n"
		}
	}else if(isAddQuestionQuery(query)){
		return "4"
	}else if(isDeleteQuestionQuery(query)){
		return "5"
	}else{
		return "1"
	}
}

// func main(){
// 	query := "equation 30/02/2023 ? Day of 4/5/2023? Day 30/02/2023?"
// 	questions := SplitQuestion(query)
// 	for i := 0; i < len(questions); i++{
// 		fmt.Println(QueryClassification(questions[i]))
// 	}
	
// }