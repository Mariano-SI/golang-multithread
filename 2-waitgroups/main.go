package main

import (
	"fmt"
	"sync"
	"time"
)

// Um WaitGroup em Go é usado para esperar que um conjunto de goroutines
// termine a execução antes de prosseguir. Ele funciona como um contador:
// cada goroutine que inicia chama Add(1) e, quando termina, chama Done().
// A goroutine principal ou outra goroutine pode chamar Wait() para bloquear
// até que o contador chegue a zero.
// É uma forma de mandar a execução do programa seguir somente quando todas
// as rotinas finalizarem, evitando a estratégia insegura de usar time.Sleep
// para aguardar goroutines.

func task(name string, wg *sync.WaitGroup) { //passamos o waitgroup como ponteiro para que seja usado nosso wg em todas as rotinas e nao uma copia independente dele
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		if name == "C"{
			time.Sleep(2 * time.Second)
		}else{
			time.Sleep(1 * time.Second)
		}
	}
	wg.Done()
}

func main() {
	myWG := sync.WaitGroup{}
	myWG.Add(3)

	go task("A", &myWG)
	go task("B", &myWG)
	go task("C", &myWG)

	myWG.Wait()
}
