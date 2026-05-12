// Channels em Go são uma forma de comunicação entre goroutines.
// Eles garantem segurança, permitindo que cada goroutine saiba exatamente
// quando pode manipular um dado compartilhado, evitando race conditions.
//
// "Don't communicate by sharing memory; share memory by communicating."
package main

import "fmt"

//goroutine 1
func main() {
	myChannel := make(chan string)

	//goroutine 2
	go func() {
		myChannel <- "Olá, Mundo" // adicionou essa string ao canal = canal cheio
	}()

	//goroutine 1

	msg := <-myChannel // canal esvaziou
	fmt.Println(msg)

	// resumo: passamos um valor da goroutine 2 pra goroutine 1
}
