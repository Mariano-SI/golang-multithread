package main

import (
	"fmt"
	"time"
)

// producer envia mensagens periodicamente.
func producer(name string, delay time.Duration) <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(delay)
			ch <- fmt.Sprintf("%s -> %d", name, i)
		}

		// Fecha o channel quando terminar.
		close(ch)
	}()

	return ch
}

func main() {
	fast := producer("FAST", 500*time.Millisecond)
	slow := producer("SLOW", 2*time.Second)

	// Timeout dispara após 6 segundos.
	timeout := time.After(6 * time.Second)

	for {

		// select espera múltiplos channels.
		select {

		// Recebe mensagens do producer rápido.
		case msg, ok := <-fast:

			// ok == false => channel fechado.
			if !ok {
				fmt.Println("FAST finished")

				// nil remove o case do select.
				fast = nil

				break
			}

			fmt.Println(msg)

		// Recebe mensagens do producer lento.
		case msg, ok := <-slow:

			if !ok {
				fmt.Println("SLOW finished")
				slow = nil
				break
			}

			fmt.Println(msg)

		// Executa caso timeout aconteça.
		case <-timeout:
			fmt.Println("TIMEOUT")
			return

		// default torna o select non-blocking.
		default:
			fmt.Println("Waiting...")
			time.Sleep(200 * time.Millisecond)
		}

		// Quando ambos forem nil, termina.
		if fast == nil && slow == nil {
			fmt.Println("All finished")
			return
		}
	}
}