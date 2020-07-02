// Package parallel executes independent tasks in parallel by using Golang fanout fanin concurrency pattern.
//
// Concurrency Made Easy
package parallel

import (
	"fmt"
	"reflect"
	"sync"
)

// Executor is an interface which has to have execute function which will be implicitly invoked to execute the concurrent tasks.
type Executor interface {
	Execute() interface{}
}

// executors is a slice of Executor
type executors []Executor

// Run starts the parallel execution of the tasks provided.
//
// It casts user-defined type tasks into `Executor interface` to builds executors.
//
// To cast the task in an `Executor interface` the user-defined type `task` needs to have method with `Execute() interface{}` signature.
func Run(tasks ...interface{}) <-chan interface{} {
	return fanin(fanout(build(tasks...))...)
}

func fanout(executors []Executor) []<-chan interface{} {
	outs := make([]<-chan interface{}, len(executors))
	for i, executor := range executors {
		outs[i] = addIn(executor)
	}
	return outs
}

func addIn(executor Executor) <-chan interface{} {
	in := make(chan interface{})
	go func(executor Executor) {
		defer close(in)
		in <- executor.Execute()
	}(executor)
	return in
}

func fanin(chans ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	out := make(chan interface{})
	wg.Add(len(chans))

	addToOut := func(c <-chan interface{}) {
		for val := range c {
			out <- val
		}
		wg.Done()
	}

	for _, c := range chans {
		go addToOut(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// Builer and Converter functions

func build(tasks ...interface{}) []Executor {
	executors := executors{}
	for _, task := range tasks {
		switch reflect.TypeOf(task).Kind() {
		case reflect.Slice:
			t := reflect.ValueOf(task)
			for i := 0; i < t.Len(); i++ {
				executors = executors.append(convert(t.Index(i).Interface()))
			}
		default:
			executors = executors.append(convert(task))
		}
	}
	return executors
}

func convert(i interface{}) (e Executor, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			fmt.Println(r)
		}
	}()
	e = i.(Executor)
	return e, err
}

func (execs executors) append(e Executor, err error) executors {
	if e != nil {
		execs = append(execs, e)
	}
	return execs
}
