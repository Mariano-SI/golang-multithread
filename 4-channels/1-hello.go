// Channels em Go são uma forma de comunicação entre goroutines.
// Eles garantem segurança, permitindo que cada goroutine saiba exatamente
// quando pode manipular um dado compartilhado, evitando race conditions.
//
// "Don't communicate by sharing memory; share memory by communicating."
package main

import "fmt"

//thread 1
func main() {
	myChannel := make(chan string)

	//thread 2
	go func() {
		myChannel <- "Olá, Mundo" // adicionou essa string ao canal = canal cheio
	}()

	//thread 1

	msg := <-myChannel // canal esvaziou
	fmt.Println(msg)

	// resumo: passamos um valor da thread 2 pra thread 1
}
