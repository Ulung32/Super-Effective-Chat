package regex

import (
	"regexp"
)

func isDateQuery(query string) bool {
	date := ".*(0[1-9]|[1-2][0-9]|3[0-1])/(0[1-9]|1[0-2])/([0-9]{4}).*"
	re := regexp.MustCompile(date)
    match := re.MatchString(query)
	return match 
}
func isMathOprQuery(query string) bool {
	mathOpr := "\\d([+|-|/|*]*\\d)*"
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
func QueryClassification (query string) int {
	if(isDateQuery(query)){
		return 3
	}else if (isMathOprQuery(query)) {
		return 2
	}else if(isAddQuestionQuery(query)){
		return 4
	}else if(isDeleteQuestionQuery(query)){
		return 5
	}else{
		return 1
	}
}
// func main(){
// 	query := "8 + 7/8+1 - 8*4"
// 	fmt.Println(QueryClassification(query))
// }