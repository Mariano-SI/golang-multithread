package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var number uint64 = 0

// Usando sync/atomic para operações atômicas.
// Atomic = Operações indivisíveis que não podem ser interrompidas.
// Diferente do mutex, não bloqueia outras goroutines; usa instruções de CPU atômicas.
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Incrementa 'number' atomicamente: thread-safe sem bloqueio.
		// Equivalente a 'number++', mas seguro para concorrência.
		atomic.AddUint64(&number, 1)

		// Lê o valor atual de forma atômica (embora não seja estritamente necessário aqui).
		current := atomic.LoadUint64(&number)

		// Responde com o número atual de visitas.
		w.Write([]byte(fmt.Sprintf("Voce teve %d visitas", current)))
	})

	// Inicia o servidor HTTP na porta 3000.
	// Cada requisição cria uma nova goroutine, mas as operações atômicas garantem segurança.
	//
	// DIFERENÇAS ENTRE MUTEX E ATOMIC:
	// - Mutex: Bloqueia o acesso (exclusão mútua), serializando operações. Mais lento,
	//   mas permite blocos complexos (ex.: ler/modificar múltiplas variáveis juntas).
	// - Atomic: Operações indivisíveis via CPU, sem bloqueio. Mais rápido para operações
	//   simples, mas limitado a tipos específicos (int, uint, etc.) e não funciona para
	//   blocos complexos.
	//
	// QUAL É MELHOR?
	// - Para operações simples (ex.: contador), atomic é melhor: mais rápido, sem overhead de bloqueio.
	// - Para operações complexas (ex.: transações), mutex é necessário para consistência.
	//
	// EXEMPLO DE TESTE: Mesmo resultado que com mutex.
	// Comando: hey -n 20000 -c 100 http://localhost:3000
	// Resultado esperado: Contador chega exatamente a 20.000 (sem race condition).
	http.ListenAndServe(":3000", nil)
}