package parallel_test

import (
	"fmt"

	"github.com/JigneshSatam/parallel"
)

type customExampleType int

// Execute -> Override `Execute() interface{}` method to make customExampleType as an Executor
func (c customExampleType) Execute() interface{} {
	return c * c
}

func Example() {
	tasks := []customExampleType{1, 2, 3, 4, 5}

	for op := range parallel.Run(tasks) {
		// Cast interface{} to desired output type
		output := op.(customExampleType)
		fmt.Println(output)
	}

	// Unordered output: 1
	// 4
	// 9
	// 16
	// 25
}
