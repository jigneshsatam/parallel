// Package parallel executes any tasks in parallel by using Golang fanout fanin concurrency pattern.
package parallel

import (
	"sync"
)

// Executor is an interface which has to have execute function which will be implicitly invoked to execute the concurrent tasks
type Executor interface {
	Execute() interface{}
}

// func BuildExecutors(tasks []interface{}) []Executor {
// 	executors := make([]Executor, len(tasks))
// 	for i, executor := range tasks {
// 		// executors[i] = append(executors, executor.(Executor))
// 		executors[i] = executor.(Executor)
// 	}
// 	return executors
// }

// Begin starts the parallel execution of the executors provided
func Begin(executors []Executor) <-chan interface{} {
	return fanin(fanout(executors)...)
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
