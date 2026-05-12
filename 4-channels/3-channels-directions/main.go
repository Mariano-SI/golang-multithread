package main

import "fmt"

func publisher(name string, ch chan<- string) { //assinatura de canal que so recebe dados
	ch <- name //canal recebe mensagem
}

func reader(ch <-chan string) { // assinatura de canal que entrega dados
	fmt.Println(<-ch)
}

// goroutine 1
func main() {
	myChannel := make(chan string)

	go publisher("Hello", myChannel)
	reader(myChannel)
}
