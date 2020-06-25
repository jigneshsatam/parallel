package parallel_test

import (
	"fmt"

	"github.com/JigneshSatam/parallel"
)

type customExampleType int

// execute -> Override `execute()` method to make customExampleType as an executor
func (c customExampleType) Execute() interface{} {
	return c * c
}

func Example() {
	inputArray := []customExampleType{1, 2, 3, 4, 5}

	// Convert input type to executor
	executorArray := make([]parallel.Executor, 0, len(inputArray))
	for _, ele := range inputArray {
		executorArray = append(executorArray, ele)
	}

	for ele := range parallel.Begin(executorArray) {
		// Cast interface{} to output type
		output := ele.(customExampleType)
		fmt.Println(output)
	}

	// Unordered output: 1
	// 4
	// 9
	// 16
	// 25
}
