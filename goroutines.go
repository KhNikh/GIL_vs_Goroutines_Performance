package main

import (
	"fmt"
	"sync"
	"time"
)

// cpuBoundTask is a CPU-bound function that increments a counter a large number of times
func cpuBoundTask(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	counter := 0
	for i := 0; i < 10000000; i++ {
		counter++
	}
}

// runTasksInGoroutines runs the CPU-bound task in multiple goroutines
func runTasksInGoroutines(numGoroutines int) {
	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go cpuBoundTask(i, &wg)
	}
	wg.Wait()
}

func main() {
	startTime := time.Now()
	runTasksInGoroutines(10) // Run the CPU-bound task in 10 goroutines
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Println("Execution time:", executionTime.Seconds(), "seconds")
}
