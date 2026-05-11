package main

import "fmt"

// Channels unbuffered: sender bloqueia até receiver ler a mensagem.
func main() {
	myChannel := make(chan int)
	go publish(myChannel)
	// consumer na thread principal (não em goroutine) para que main não termine
	// antes de publish e consumer trocarem dados
	consumer(myChannel)
}

// Lê continuamente do channel até ser fechado
func consumer(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
	fmt.Println("Channel fechado, consumer terminado")
}

// Envia valores e fecha o channel
func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		// Envia e bloqueia até consumer ler
		ch <- i
	}
	close(ch)
}