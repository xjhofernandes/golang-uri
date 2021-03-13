// Entrada

// A primeira linha de entrada contém um único inteiro positivo N (1 < N < 105) Este é o número de linhas de entrada que vem logo a seguir. As próximas N linhas conterão, cada uma delas, um valor inteiro não negativo.
// Saída

// Apresente todos os valores lidos na entrada segundo a ordem apresentada acima. Cada número deve ser impresso em uma linha, conforme exemplo abaixo.

package main

import (
	"fmt"
	"sort"
)

func main() {
	var entradas int

	fmt.Scan(&entradas)

	var numerosPares []int
	var numerosImpares []int

	for i := 0; i < entradas; i++ {
		var numero int

		fmt.Scan(&numero)

		if numero%2 == 0 && ChecarSeExiste(numero, numerosPares) {
			numerosPares = append(numerosPares, numero)
		} else if ChecarSeExiste(numero, numerosImpares) {
			numerosImpares = append(numerosImpares, numero)
		}
	}

	sort.Ints(numerosPares)
	numerosImpares = CriarInverso(numerosImpares)

	ImprimirNumeros(numerosPares)
	ImprimirNumeros(numerosImpares)
}

func ChecarSeExiste(numero int, arrayNumeros []int) bool {
	for _, numeroArray := range arrayNumeros {
		if numero == numeroArray {
			return false
		}
	}
	return true
}

func CriarInverso(arrayNumeros []int) []int {
	sort.Ints(arrayNumeros)
	var arrayInvertido []int

	for i := len(arrayNumeros) - 1; i >= 0; i-- {
		arrayInvertido = append(arrayInvertido, arrayNumeros[i])
	}

	return arrayInvertido
}

func ImprimirNumeros(arrayNumeros []int) {
	for _, numeroArray := range arrayNumeros {
		fmt.Println(numeroArray)
	}
}
