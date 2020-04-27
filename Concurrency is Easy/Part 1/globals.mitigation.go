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
// ➜ git:(master) ✗ go build -o no-globals global-state.mitigation.go
// ➜ git:(master) ✗ ./no-globals  
//  [worker 4] counter: 0
//  [worker 4] counter: 1
//  [worker 4] counter: 2
//  [worker 4] counter: 3
//  [worker 4] counter: 4
//  [worker 3] counter: 0
//  [worker 3] counter: 1
//  [worker 3] counter: 2
//  [worker 3] counter: 3
//  [worker 3] counter: 4
//  [worker 0] counter: 0
//  [worker 0] counter: 1
//  [worker 2] counter: 0
//  [worker 2] counter: 1
//  [worker 2] counter: 2
//  [worker 1] counter: 0
//  [worker 1] counter: 1
//  [worker 1] counter: 2
//  [worker 1] counter: 3
//  [worker 1] counter: 4
//  [worker 0] counter: 2
//  [worker 0] counter: 3
//  [worker 0] counter: 4
//  [worker 2] counter: 3
//  [worker 2] counter: 4

package main

import (
	"sync"
	"fmt"
)

func main() {
	wg := sync.WaitGroup{}

	for i:=0; i<5; i++ {
		fmt.Println(i)
		go incrementer()(&wg, i)
	}

	wg.Wait()
}

func incrementer() func(*sync.WaitGroup, int) {
	counter := 0

	return func (wg *sync.WaitGroup, workerId int) {
		wg.Add(1)
	
		for i:=0; i<5; i++ {
			fmt.Printf("[worker %d] counter: %d\n", workerId, counter)
			counter++
		}
	
		wg.Done()
	}
}

