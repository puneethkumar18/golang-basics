package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, wg *sync.WaitGroup) {
	fmt.Println("worker", i, "started")
	fmt.Println("worker", i, "Ended")
	wg.Done()
}

func main() {
	fmt.Println("Learning Sync ....")

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	//wg.Wait()
	time.Sleep(time.Microsecond * 1000)
	fmt.Println("Workers complted")
	//time.Sleep(time.Millisecond * 10)
}
