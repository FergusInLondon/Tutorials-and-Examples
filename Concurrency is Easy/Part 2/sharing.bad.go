//
//
//
//
//
//
//
//
//
//
//
// ➜ git:(master) ✗ go build -o bad-globals global-state.bad.go
// ➜ git:(master) ✗ ./bad-globals  
//  [worker 4] counter: 0
//  [worker 4] counter: 1
//  [worker 4] counter: 2
//  [worker 4] counter: 3
//  [worker 4] counter: 4
//  [worker 1] counter: 5
//  [worker 1] counter: 6
//  [worker 1] counter: 7
//  [worker 1] counter: 8
//  [worker 1] counter: 9
//  [worker 0] counter: 0
//  [worker 0] counter: 11
//  [worker 0] counter: 12
//  [worker 0] counter: 13
//  [worker 2] counter: 5
//  [worker 0] counter: 14
//  [worker 2] counter: 16
//  [worker 2] counter: 17
//  [worker 2] counter: 18
//  [worker 2] counter: 19
//  [worker 3] counter: 10
//  [worker 3] counter: 21
//  [worker 3] counter: 22
//  [worker 3] counter: 23
//  [worker 3] counter: 24
package main

import (
	"sync"
	"fmt"
)

func main() {
	wg := sync.WaitGroup{}

	for i:=0; i<5; i++ {
		go incrementer(&wg, i)
	}

	wg.Wait()
}

var counter = 0

func incrementer(wg *sync.WaitGroup, workerId int) {
	wg.Add(1)

	for i:=0; i<5; i++ {
		fmt.Printf("[worker %d] counter: %d\n", workerId, counter)
		counter++
	}

	wg.Done()
}