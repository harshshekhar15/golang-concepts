// Program to solve race condition using channels --
// Here we launch 1000 go routines that try to update counter variable
// and using channels we make sure that at a given point of time only
// go routines updates the counter
package channel

import (
	"fmt"
	"sync"
)

func RunRaceProgram() {
	fmt.Println("Running program to solve race condition using channels -->")

	var counter int = 0
	var wg sync.WaitGroup
	lock := make(chan bool, 1)

	// launch 1000 go routines
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, lock chan bool) {
			// send a value to the lock channel so that
			// other go routines are blocked
			lock <- true
			// update counter
			counter++
			<-lock
			wg.Done()
		}(&wg, lock)
	}
	wg.Wait()
	fmt.Println("Count:", counter)
}
