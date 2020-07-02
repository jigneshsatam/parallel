package parallel_test

import (
	"fmt"

	"github.com/JigneshSatam/parallel"
)

// customType -> Let `customType` be a user-defined type
type customType int

// Execute -> Create `Execute() interface{}` method for customType
func (c customType) Execute() interface{} {
	return c * c
}

func Example() {
	tasks := []customType{1, 2, 3, 4, 5}

	// Call `parallel.Run()` to start parallel execution
	outputChannel := parallel.Run(tasks)

	for op := range outputChannel {
		// Cast interface{} to desired output type
		output := op.(customType)
		fmt.Println(output)
	}

	// Unordered output: 1
	// 4
	// 9
	// 16
	// 25
}
