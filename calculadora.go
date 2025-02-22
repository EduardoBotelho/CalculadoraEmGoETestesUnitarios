package main

import (
	"fmt"
	"os"
)

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()
		switch comando {
		case 1:
			x := soma(2, 3)
			fmt.Println("Resultado da Soma:", x)
		case 2:
			y := subtracao(3, 2)
			fmt.Println("Resultado da Subtração:", y)
		case 3:
			z := multiplica(2, 5)
			fmt.Println("Resultado da Multiplicação:", z)
		case 4:
			w := divisao(6, 3)
			fmt.Println("Resultado da Divisão:", w)
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido")
			os.Exit(-1)
		}
	}
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)
	return comandoLido
}

func exibeIntroducao() {
	fmt.Println("----Calculadora----")
	fmt.Println("Escolha uma opção")
}

func exibeMenu() {
	fmt.Println("Escolha uma Opção:")
	fmt.Println("1 - Soma")
	fmt.Println("2 - Subtração")
	fmt.Println("3 - Multiplicação")
	fmt.Println("4 - Divisão")
	fmt.Println("0 - Sair")
}

func soma(numeros ...int) int {
	total := 0
	for _, v := range numeros {
		total += v
	}
	return total
}

func subtracao(numeros ...int) int {
	if len(numeros) == 0 {
		return 0
	}
	total := numeros[0]
	for _, v := range numeros[1:] {
		total -= v
	}
	return total
}

func multiplica(numeros ...int) int {
	if len(numeros) == 0 {
		return 0
	}
	total := 1
	for _, v := range numeros {
		total *= v
	}
	return total
}

func divisao(numeros ...int) float64 {
	if len(numeros) == 0 {
		return 0
	}
	total := float64(numeros[0])
	for _, v := range numeros[1:] {
		if v == 0 { //Verificacao para tratar divisao por zero
			fmt.Println("Erro: Divisão por zero")
			return 0
		}
		total /= float64(v)
	}
	return total
}
