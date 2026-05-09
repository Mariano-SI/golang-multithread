package main

import (
	"fmt"
	"net/http"
)

var number uint64 = 0

func main() {
	// AVISO: Este exemplo está suscetível a race condition!
	//
	// O servidor http.ListenAndServe cria uma nova goroutine para cada requisição.
	// Quando múltiplas requisições chegam simultaneamente, múltiplas goroutines
	// tentam modificar a variável 'number' ao mesmo tempo sem sincronização.
	// Isso pode levar a race condition: incrementos perdidos, valores inconsistentes,
	// ou comportamento não determinístico.
	//
	// Para demonstrar, simulamos 10.000 requisições simultâneas usando o comando:
	// hey -n 10000 -c 50 http://localhost:3000
	//
	// Em nossos testes, após as 10.000 requisições, o contador 'number' ficou em 9.365
	// em vez de 10.000. Isso aconteceu porque múltiplas goroutines acessaram e
	// incrementaram a variável ao mesmo tempo, causando perda de incrementos devido
	// à race condition.
	//
	// Para evitar isso, seria necessário usar mutex (sync.Mutex) ou atomic operations.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		w.Write([]byte(fmt.Sprintf("Voce teve %d visitas", number)))
	})
	http.ListenAndServe(":3000", nil)
}
