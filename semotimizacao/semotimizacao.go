package main

import (
	"fmt"
	"time"
)

func verificaSeNumeroEhPrimo(num int) {

	var primo bool = true

	if num < 2 {
		primo = false
	}

	for i := 2; i < num; i++ {

		if num%i == 0 && num != i {
			primo = false
		}

	}

	fmt.Printf("o número %d é primo? %t\n", num, primo)
	fmt.Printf("Fim do número %d: %s\n", num, time.Now().Format("01-02-2006 15:04:05"))

}

func verificaSeNumerosSaoPrimos(num []int) {

	for i := 0; i < len(num); i++ {
		verificaSeNumeroEhPrimo(num[i])
	}

}

func main() {

	inicio := time.Now()

	var numeros []int
	numeros = append(numeros, 100000000000)
	numeros = append(numeros, 100000000000)

	verificaSeNumerosSaoPrimos(numeros)

	fmt.Println("\nInicio do Programa:", inicio.Format("01-02-2006 15:04:05"))
	fmt.Println("Fim do Programa:", time.Now().Format("01-02-2006 15:04:05"))
}
