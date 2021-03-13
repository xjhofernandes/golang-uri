package main

import (
	"fmt"
)

func main() {
	var entradas int

	fmt.Scan(&entradas)

	var numeros []int

	for i := 0; i < entradas; i++ {
		var numeroArray int

		fmt.Scan(&numeroArray)
		numeros = append(numeros, numeroArray)
	}

	numeros = OrdenarNumeros(numeros, entradas)

	fmt.Print(numeros)
}

func OrdenarNumeros(numerosArray []int, tamanhoArray int) []int {
	var temp int

	for y := 0; y < tamanhoArray-1; y++ {
		for i := 0; i < tamanhoArray-1; i++ {
			if numerosArray[i] > numerosArray[i+1] {
				temp = numerosArray[i+1]
				numerosArray[i+1] = numerosArray[i]
				numerosArray[i] = temp
			}
		}
	}

	return numerosArray
}
