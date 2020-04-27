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
// ➜ git:(master) ✗ go build -o mux-sharing sharing.mitigation..go
// ➜ git:(master) ✗ ./mux-sharing
// [worker 0] counter: 0
// [worker 0] counter: 1
// [worker 0] counter: 2
// [worker 0] counter: 3
// [worker 0] counter: 4
// [worker 4] counter: 5
// [worker 4] counter: 6
// [worker 4] counter: 7
// [worker 4] counter: 8
// [worker 4] counter: 9
// [worker 2] counter: 10
// [worker 2] counter: 11
// [worker 2] counter: 12
// [worker 2] counter: 13
// [worker 2] counter: 14
// [worker 3] counter: 15
// [worker 3] counter: 16
// [worker 3] counter: 17
// [worker 3] counter: 18
// [worker 3] counter: 19
// [worker 1] counter: 20
// [worker 1] counter: 21
// [worker 1] counter: 22
// [worker 1] counter: 23
// [worker 1] counter: 24

package main

import (
	"sync"
	"fmt"
)

func main() {
	wg := sync.WaitGroup{}
	mux := sync.Mutex{}

	for i:=0; i<5; i++ {
		go incrementer(&wg, &mux, i)
	}

	wg.Wait()
}

var counter = 0



func incrementer(wg *sync.WaitGroup, mux *sync.Mutex, workerID int) {
	wg.Add(1)

	for i:=0; i<5; i++ {
		mux.Lock()

		fmt.Printf("// [worker %d] counter: %d\n", workerID, counter)
		counter++

		mux.Unlock()
	}

	wg.Done()
}