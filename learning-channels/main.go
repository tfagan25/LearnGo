package main

import (
	"fmt"
	"time"
	"sync"
)

type Whatever struct {
	OutputMessage string
	Seconds int
}

func waitTime(ch <-chan Whatever, wg *sync.WaitGroup) {
	item := <- ch

	time.Sleep(time.Duration(int(time.Second) * item.Seconds))
	fmt.Println(item.OutputMessage)
	wg.Done()
}

func main () {
	var wg sync.WaitGroup
	mail := make(chan Whatever)

	first := Whatever{"first one", 1}
	second := Whatever{"second one", 2}
	third := Whatever{"third one", 3}

	wg.Add(3)

	go waitTime(mail, &wg)
	go waitTime(mail, &wg)
	go waitTime(mail, &wg)

	mail <- first
	mail <- second
	mail <- third

	wg.Wait()
}