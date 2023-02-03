package parallel_test

import (
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

var testCases = []struct {
	name string
	ip   interface{}
	op   []int
}{
	{
		name: "Valid Tasks",
		ip:   []customType{1, 2, 3, 4, 5},
		op:   []int{1, 4, 9, 16, 25},
	},
	{
		name: "Invalid Tasks - where Execute method not overridden by input type",
		ip:   []int{1, 5},
		op:   []int{},
	},
	{
		name: "Valid Single Task",
		ip:   customType(5),
		op:   []int{25},
	},
	{
		name: "Inalid Single Task - where Execute method not overridden by input type",
		ip:   5,
		op:   []int{},
	},
	{
		name: "Valid Empty Tasks",
		ip:   []customType{},
		op:   []int{},
	},
	{
		name: "Invalid Empty Tasks - where Execute method not overridden by input type",
		ip:   []int{},
		op:   []int{},
	},
}

func TestRun(t *testing.T) {
	t.Parallel()
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			outputArray := []customType{}
			for ele := range parallel.Run(tc.ip) {
				outputArray = append(outputArray, ele.(customType))
			}

			intArr := []int{}
			for _, ele := range outputArray {
				intArr = append(intArr, int(ele))
			}

			sort.Ints(intArr)

			if !reflect.DeepEqual(tc.op, intArr) {
				t.Errorf("Run Failed for test %v \n want ===> %v \n got  ===> %v", tc.name, tc.op, intArr)
			}
		})
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []customType{1, 2, 3, 4, 5}
		op := []int{1, 4, 9, 16, 25}
		output := []customType{}
		for ele := range parallel.Run(arr) {
			output = append(output, ele.(customType))
		}

		intArr := []int{}
		for _, ele := range output {
			intArr = append(intArr, int(ele))
		}

		sort.Ints(intArr)

		if !reflect.DeepEqual(op, intArr) {
			b.Errorf("Run Failed for test \n want ===> %v \n got  ===> %v", op, intArr)
		}

	}
}
