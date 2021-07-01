// Simple mutex program that creates 1000 go routines which in turn try to update
// a counter variable by taking a lock on it and unlocking it once it is done
package mutex

import (
	"fmt"
	"sync"
)

func Run() {
	fmt.Println("Running mutex program")

	var counter int = 0
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
			mutex.Lock()
			counter++
			mutex.Unlock()
			wg.Done()
		}(&wg, mutex)
	}

	wg.Wait()
	fmt.Println("Count:", counter)
}
