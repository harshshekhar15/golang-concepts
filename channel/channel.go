package channel

import (
	"fmt"
	"math/rand"
	"sync"
)

func RunProgram1() {
	fmt.Println("Starting program. Hold tight!!")

	// Create a channel for sending and receiving job objects
	jobChannel := make(chan Job, 1500)

	var wg sync.WaitGroup
	wg.Add(3)

	// Start producer go routines
	go producer("producer1", jobChannel, &wg)
	go producer("producer2", jobChannel, &wg)
	go producer("producer3", jobChannel, &wg)

	// Wait for the producers to complete their work
	// and close the jobChannel once they are done
	wg.Wait()
	close(jobChannel)

	// done will be used to track the status of
	// consumer go routine
	done := make(chan bool)

	// Start consumer go routine
	go consumer(jobChannel, done)
	<-done
}

type Job struct {
	num  int
	Name string
}

func producer(name string, ch chan<- Job, wg *sync.WaitGroup) {
	itemsToProduce := rand.Intn(500)
	fmt.Println("Items Produced by", name, ":", itemsToProduce)
	for i := 0; i < itemsToProduce; i++ {
		ch <- Job{
			num:  rand.Intn(10),
			Name: name,
		}
	}
	wg.Done()
}

func consumer(ch <-chan Job, done chan<- bool) {
	var count int = 0
	for job := range ch {
		fmt.Println(job.Name, ":", job.num)
		count++
	}
	fmt.Println("Total items produced:", count)
	done <- true
}
