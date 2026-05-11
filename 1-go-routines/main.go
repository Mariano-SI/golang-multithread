package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := range 10 {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1
func main() {
	//Sem routines = execucao sequencial
	initialTimeWithoutRoutines := time.Now()

	task("A")
	task("B")

	finalTimeWithoutRoutines := time.Now()

	totalTimeWithoutRoutines := finalTimeWithoutRoutines.Sub(initialTimeWithoutRoutines)

	fmt.Printf("Tempo total de execução: %v\n", totalTimeWithoutRoutines) // 20 sec. Executou A depois B

	// Com routines
	// Aqui criamos duas goroutines com a palavra-chave go.
	// Isso introduz concorrência: as duas tarefas podem progredir 'ao mesmo tempo'
	// do ponto de vista do programa. Se a máquina tiver múltiplos núcleos, o
	// runtime do Go pode até executar ambas em paralelo verdadeiro.
	// No entanto, o modelo de goroutines em Go é fundamentalmente concorrente;
	// o paralelismo depende do agendamento e da disponibilidade de múltiplos CPUs.

	//thread 2
	go task("A")
	//thread 3
	go task("B")

	//thread 4
	go func() {
		task("C")
	}()

	fmt.Println("As goroutines são iniciadas e podem rodar concorrentemente/parallelamente")

	// Sem esse sleep, a função main terminaria e o programa encerraria antes
	// que as goroutines terminassem. Usamos time.Sleep apenas para manter o
	// processo vivo tempo suficiente para a saída aparecer no terminal.
	time.Sleep(15 * time.Second) // Mantém o programa aberto para ver a saída das goroutines

}
