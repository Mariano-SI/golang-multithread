package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Existem duas formas comuns de usar WaitGroup com goroutines:
	//
	// 1. Alterar a assinatura da função para receber *sync.WaitGroup como parâmetro
	//    (visto em main.go)
	//
	// 2. Manter a assinatura da função normal e usar uma função anônima (closure)
	//    para envolver a execução e passar o WaitGroup (visto aqui)
	//
	// A primeira abordagem é mais comum em Go, pois é explícita e clara.
	// A segunda abordagem permite manter a função simples sem acoplamento ao WaitGroup,
	// o que pode ser útil se a função precisar ser usada em diferentes contextos.

	var wg sync.WaitGroup

	wg.Add(3)

	go func(name string) {
		defer wg.Done()
		task(name)
	}("A")

	go func(name string) {
		defer wg.Done()
		task(name)
	}("B")

	go func(name string) {
		defer wg.Done()
		task(name)
	}("C")

	wg.Wait()
	fmt.Println("Todas as goroutines terminaram")
}
