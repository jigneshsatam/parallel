# parallel
### **Concurrency Made Easy**
***
### [![Go Report Card](https://goreportcard.com/badge/github.com/JigneshSatam/parallel)](https://goreportcard.com/report/github.com/JigneshSatam/parallel) [![Build](https://github.com/JigneshSatam/parallel/actions/workflows/go.yml/badge.svg)](https://github.com/JigneshSatam/parallel/actions/workflows/go.yml) <!--[![Build Status](https://api.travis-ci.org/JigneshSatam/parallel.svg?branch=master)](https://travis-ci.org/JigneshSatam/parallel)--> [![Test Coverage](https://api.codeclimate.com/v1/badges/20812caef2634511a1e6/test_coverage)](https://codeclimate.com/github/JigneshSatam/parallel/test_coverage) [![Issue Count](https://codeclimate.com/github/JigneshSatam/parallel/badges/issue_count.svg)](https://codeclimate.com/github/JigneshSatam/parallel) [![Code Climate](https://codeclimate.com/github/JigneshSatam/parallel/badges/gpa.svg)](https://codeclimate.com/github/JigneshSatam/parallel) [![GoDoc](https://godoc.org/github.com/JigneshSatam/parallel?status.svg)](https://godoc.org/github.com/JigneshSatam/parallel) <!--[![Sourcegraph](https://sourcegraph.com/github.com/JigneshSatam/parallel/-/badge.svg)](https://sourcegraph.com/github.com/JigneshSatam/parallel?badge)--> [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/JigneshSatam/parallel/blob/master/LICENSE)

### **parallel executes independent tasks in parallel by using Golang fanout fanin concurrency pattern**

## Installation

```go
go get -u github.com/jigneshsatam/parallel
```

## Usage

Example code snippet: https://go.dev/play/p/Hf6jkXNFB6E

### Step 1: Import the package
```go
import (
	"github.com/jigneshsatam/parallel"
)
```

### Step 2: Create `Execute() interface{}` method for `user defined type`

- Let `employee` be a user-defined type
- Suppose fetching `employee details` is an independent task and needs to be  done be in parallel
- Create a method `Execute()` for `employee` type
- Signature of Execute method `Execute() interface{}`
- Add the code to fetch the `employee details` in this `Execute()` method
```go
// employee -> Let `employee` be a user-defined type
type employee struct {
	name string
}

// Execute() -> Create `Execute() interface{}` method for employee type
func (e employee) Execute() interface{} {
	return employeeDetails(e)
}

func employeeDetails(e employee) string {
	return "Employee details Name: " + e.name
}
```

### Step 3: Call `parallel.Run()` method

- Call `parallel.Run` by passing list of employees who's details to be fetched.
```go
func Example() {
	employees := []employee{employee{"foo"}, employee{"bar"}, employee{"baz"}}

	// Call `parallel.Run()` to start parallel execution
	outputChannel := parallel.Run(employees)

	for op := range outputChannel {
		// Cast `interface{}` to desired output type
		output := op.(string)
		fmt.Println(output)
	}
}

// Unordered output:
// "Employee details Name: bar"
// "Employee details Name: foo"
// "Employee details Name: baz"
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update the tests as appropriate.

## License
This project is licensed under the **MIT** License - see the [LICENSE.md](https://github.com/jigneshsatam/parallel/blob/master/LICENSE) file for details
