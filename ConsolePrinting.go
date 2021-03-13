package main

import (
	"fmt"
)

type ContaBancaria struct {
	nome          string
	idade         int
	saldo         float64
	cadastroAtivo bool
}

func main() {

	var numero int = 66
	var palavra string = "'Pequeno Teste'"
	var saida string
	var valor float64 = 1579.5789

	var boleano bool = true

	//fmt.Scan(&numero)
	//fmt.Scan(&palavra)

	saida = fmt.Sprintf("palavra: %s e número: %d", palavra, numero)
	fmt.Println(saida)

	saida = fmt.Sprintf("Valor Boleano: %t", boleano)
	fmt.Println(saida)

	saida = fmt.Sprintf("Valor Boleano: %.4f", valor)
	fmt.Println(saida)

	var conta ContaBancaria
	conta = ContaBancaria{nome: "John Johns", idade: 24, saldo: 6879.787845, cadastroAtivo: true}

	fmt.Println(conta)
	fmt.Println("Texto formatado:")
	fmt.Println("==============================")
	saida = fmt.Sprintf("Nome: %s\nidade: %d\nsaldo em conta: %f\nCadastro ativo: %t", conta.nome, conta.idade, conta.saldo, conta.cadastroAtivo)
	fmt.Println(saida)
	fmt.Println("==============================")
}
