# parallel [![Go Report Card](https://goreportcard.com/badge/github.com/JigneshSatam/parallel)](https://goreportcard.com/report/github.com/JigneshSatam/parallel) [![Build Status](https://travis-ci.org/JigneshSatam/parallel.svg?branch=master)](https://travis-ci.org/JigneshSatam/parallel) [![Test Coverage](https://api.codeclimate.com/v1/badges/20812caef2634511a1e6/test_coverage)](https://codeclimate.com/github/JigneshSatam/parallel/test_coverage) [![Issue Count](https://codeclimate.com/github/JigneshSatam/parallel/badges/issue_count.svg)](https://codeclimate.com/github/JigneshSatam/parallel) [![Code Climate](https://codeclimate.com/github/JigneshSatam/parallel/badges/gpa.svg)](https://codeclimate.com/github/JigneshSatam/parallel) [![GoDoc](https://godoc.org/github.com/JigneshSatam/parallel?status.svg)](https://godoc.org/github.com/JigneshSatam/parallel) [![Sourcegraph](https://sourcegraph.com/github.com/JigneshSatam/parallel/-/badge.svg)](https://sourcegraph.com/github.com/JigneshSatam/parallel?badge) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/JigneshSatam/parallel/blob/master/LICENSE)

parallel executes any tasks in parallel by using Golang fanout fanin concurrency pattern

## Installation

```go
go get -u github.com/JigneshSatam/parallel
```

## Usage

### Step 1: Import the package
```go
import (
  "github.com/JigneshSatam/parallel"
)
```

### Step 2: Override  `Execute() interface{}` method for `user defined type`

- Let `user` be a user defined type.
- Suppose fetching `user's details` needs to done be in parallel.
- Then override `Execute() interface{}` method to make `user` as an `Executor`.
- And write the code to fetch the `user details` in this `Execute()` method.
```go
// user -> Let `user` be a user defined type
type user struct {
  name string
}

// Execute() -> Overriden `Execute() interface{}` method
func (u user) Execute() interface{} {
  return "Name: " + u.name + " and other details..."
}
```

### Step 3: Call `parallel.Run()` method

- Call `parallel.Run` by passing list of users who's details to be fetched.
```go
func Example() {
  users := []user{user{"foo"}, user{"bar"}, user{"baz"}}

  outputChannel :=  parallel.Run(users)

  for op := range outputChannel {
    // Cast `interface{}` to desired output type
    output := op.(string)
    fmt.Println(output)
  }
}

// Unordered output:
// "Name: bar and other details..."
// "Name: foo and other details..."
// "Name: baz and other details..."
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update the tests as appropriate.

## License
This project is licensed under the **MIT** License - see the [LICENSE.md](https://github.com/JigneshSatam/parallel/blob/master/LICENSE) file for details
