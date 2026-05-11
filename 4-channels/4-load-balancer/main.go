package main

import (
	"fmt"
	"sync"
	"time"
)


// worker processa os valores recebidos do channel.
//
// Cada valor enviado no channel será recebido
// por apenas UM worker.
//
// Isso evita acesso concorrente ao mesmo dado,
// sem necessidade de mutex.
func worker(workerId int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}

	fmt.Printf("Worker %d finished\n", workerId)
}

func main() {
	// Channel usado para distribuir jobs
	data := make(chan int)

	workersQuantity := 10

	var wg sync.WaitGroup

	for i := range workersQuantity {
		wg.Add(1)
		go worker(i, data, &wg)
	}

	// Envia jobs para o channel
	for i := range 100 {
		data <- i
	}

	// Fecha o channel:
	// avisa aos workers que não existem mais jobs
	close(data)

	// Espera todos os workers terminarem
	wg.Wait()

	fmt.Println("All workers finished")
}


/*
EXPLICAÇÃO DE PERFORMANCE

Cada tarefa demora 1 segundo:

	time.Sleep(time.Second)

Se executássemos 100 tarefas sequencialmente:

	100 tarefas × 1 segundo = ~100 segundos

Com 10 workers concorrentes:

	100 tarefas / 10 workers = ~10 segundos

Os workers executam tarefas em paralelo,
reduzindo drasticamente o tempo total.

Esse padrão é chamado de "worker pool".

O channel sincroniza automaticamente as goroutines:

	data <- i

O envio bloqueia até algum worker receber o valor.

Como cada valor é entregue para apenas um worker,
não existe concorrência sobre a mesma tarefa.
*/
