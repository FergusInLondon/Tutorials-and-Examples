/*
➜  Part 2 git:(master) ✗ go build -o once once.bad.go
➜  Part 2 git:(master) ✗ ./once                      
Initialising!
Initialising!
Initialising!
➜  Part 2 git:(master) ✗ ./once
Initialising!
Initialising!
➜  Part 2 git:(master) ✗ ./once
Initialising!
Initialising!
➜  Part 2 git:(master) ✗ ./once
Initialising!
Initialising!
Initialising!
Initialising!
Initialising!
Initialising!
*/

package main

import (
	"sync"
	"fmt"
)

var initialised = false
func doAction() {
	fmt.Println("Initialising!")
	initialised = true
}

func main() {
	wg := &sync.WaitGroup{}
	for i:=0; i<50; i++ {
		go func(){
			wg.Add(1)
			for i:=0; i<50; i++ {
				if !initialised {
					doAction()
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
}%           