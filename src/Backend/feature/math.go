package feature

import (
    "errors"
    "github.com/Knetic/govaluate"
    "math"
	// "fmt"
)

func MathematicalOperationSolver(mathematicalExpression string) (float64, error) {
    expression, err := govaluate.NewEvaluableExpression(mathematicalExpression)
    if err != nil {
        return 0, errors.New("Invalid Expression")
    }
    result, err := expression.Evaluate(nil)
    if err != nil {
        return 0, errors.New("Invalid Expression")
    }
    if val, ok := result.(float64); ok {
        if math.IsInf(val, 0) {
            return 0, errors.New("Division by zero")
        }
    }
    return result.(float64), nil
}

// func main() {
// 	expression := "1+2+(3*5)/0"
// 	result, status := MathematicalOperationSolver(expression)
// 	if status != nil {
// 		fmt.Println(status)
// 	} else {
// 		fmt.Println("This is the result", result)
// 	}
// }