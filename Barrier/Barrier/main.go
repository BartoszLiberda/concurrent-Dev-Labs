// Author: Bartosz Liberda
// Student Number: C00295791

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, wg2 *sync.WaitGroup) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	wg2.Done() // decrement the new waitgroup by 1 once Part A has ran
	wg2.Wait() // wait until the waigroup counter is at 0 and then proceed to run Part B
	fmt.Println("PartB", goNum)
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup // initialised new waitgroup to track Part A
	wg.Add(totalRoutines)
	wg2.Add(totalRoutines) // Set new waitgroup to 10
	//we will need some of these
	ctx := context.TODO()
	var theLock sync.Mutex
	sem := semaphore.NewWeighted(int64(totalRoutines))
	theLock.Lock()
	sem.Acquire(ctx, 1)
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &wg, &wg2) // pass the new waitgroup into the function
	}
	sem.Release(1)
	theLock.Unlock()

	wg.Wait() //wait for everyone to finish before exiting
}

// People I Helped: Ciaran O'Toole, Samuel Geraghty
