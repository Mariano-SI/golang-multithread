package main

import (
	"fmt"
	"time"
)

func main() {

	// Buffer de 3 posições.
	ch := make(chan string, 3)

	/*
		Sem buffer:

			make(chan string)

		o sender bloqueia após cada envio,
		esperando alguém consumir.

		Com buffer:
		o channel acumula mensagens temporariamente.
	*/

	go func() {

		fmt.Println("Sending 1")
		ch <- "msg 1"

		fmt.Println("Sending 2")
		ch <- "msg 2"

		fmt.Println("Sending 3")
		ch <- "msg 3"

		// Aqui o buffer encheu.
		fmt.Println("Sending 4 (waiting)")
		ch <- "msg 4"

		fmt.Println("Sent 4")

		close(ch)
	}()

	// Espera para mostrar buffer enchendo.
	time.Sleep(2 * time.Second)

	for msg := range ch {
		fmt.Println("Received:", msg)
		time.Sleep(time.Second)
	}
}