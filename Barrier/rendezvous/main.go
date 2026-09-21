// Author: Bartosz Liberda
// Student Number: C00295791

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

//Global variables shared between functions --A BAD IDEA

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, barrier chan bool) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	<-barrier
	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	barrier := make(chan bool)
	threadCount := 10

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, barrier)
		barrier <- true
	}
	wg.Wait() //wait here until everyone (10 go routines) is done

}
