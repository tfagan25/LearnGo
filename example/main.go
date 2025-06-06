package main

import (
	"time"
	"sync"
)

func wait(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(10 * time.Second)
}

func main () {
	var wg sync.WaitGroup

	wg.Add(1)
	go wait(&wg)
	wg.Add(1)
	go wait(&wg)

	wg.Wait()
}