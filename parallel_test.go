package parallel_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/JigneshSatam/parallel"
)

type customType int

// execute -> Overridden `execute()` method to make customType as an executor
func (c customType) Execute() interface{} {
	return c * c
}

func TestBegin(t *testing.T) {
	t.Parallel()
	inputArray := []customType{1, 2, 3, 4, 5}
	want := []int{1, 4, 9, 16, 25}
	executorArray := make([]parallel.Executor, 0, len(inputArray))

	// Convert input to executor
	for _, ele := range inputArray {
		executorArray = append(executorArray, ele)
	}

	outputArray := []customType{}
	for ele := range parallel.Begin(executorArray) {
		outputArray = append(outputArray, ele.(customType))
	}

	intArr := []int{}
	for _, ele := range outputArray {
		intArr = append(intArr, int(ele))
	}

	sort.Ints(intArr)

	if !reflect.DeepEqual(want, intArr) {
		t.Errorf("Go Failed \n want ===> %v \n got  ===> %v", want, intArr)
	}
}

func BenchmarkBegin(b *testing.B) {
	for i := 0; i < b.N; i++ {

		// }
		arr := []customType{1, 2, 3, 4, 5}
		arrN := make([]parallel.Executor, 0, len(arr))
		for _, ele := range arr {
			arrN = append(arrN, ele)
		}
		x := parallel.Begin(arrN)
		y := make([]customType, 0, len(arr))
		// fmt.Println(arrN)
		for ele := range x {
			y = append(y, ele.(customType))
		}
		// fmt.Println(y)
	}
}

func ExampleBegin() {
	inputArray := []customType{1, 2, 3, 4, 5}

	// Convert input type to executor
	executorArray := make([]parallel.Executor, 0, len(inputArray))
	for _, ele := range inputArray {
		executorArray = append(executorArray, ele)
	}

	for ele := range parallel.Begin(executorArray) {
		// Cast interface{} to output type
		output := ele.(customType)
		fmt.Println(output)
	}

	// Unordered output: 1
	// 4
	// 9
	// 16
	// 25
}
