package feature

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"errors"
	"log"
	// "math"
	// "strings"
	// "strconv"
)

func MathematicalOperationSolver(mathematicalExpression string) (float64, error) {
	expression, err := govaluate.NewEvaluableExpression(mathematicalExpression)
	if(err != nil){
		return 0, errors.New("Invalid Expression")
	}
	result, err := expression.Evaluate(nil)
	if(err != nil){
		return 0, errors.New("Invalid Expression")
	}
	return result.(float64), nil
}

// func main(){
// 	expression := "1+2+((3*5)"
// 	result, status := MathematicalOperationSolver(expression)
// 	if(status != nil){
// 		log.Fatal(status)
// 	}
// 	fmt.Println("This is the result", result)
// }