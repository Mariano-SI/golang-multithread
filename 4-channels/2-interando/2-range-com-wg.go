package main

import (
	"fmt"
	"sync"
)

func main() {
	myChannel := make(chan int)
	myWg := sync.WaitGroup{}

	myWg.Add(2)

	go func() {
		defer myWg.Done()
		publish(myChannel)
	}()

	go func() {
		defer myWg.Done()
		consumer(myChannel)
	}()

	myWg.Wait()
}

func consumer(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
	fmt.Println("Channel fechado, consumer terminado")
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		
		ch <- i
	}
	close(ch)
}
