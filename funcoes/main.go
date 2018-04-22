package main

import (
	"fmt"
)

func soma(x int, y int) int {

	return x + y

}

func troca(x string, y string) (string, string) {
	return y, x
}

func divide(x, y int) (resultado, resto int) {

	resto = x % y

	resultado = x / y

	return
}

func executaFuncao(f func(string) string, valor string) {

	aux := f(valor)

	fmt.Printf(aux)
}

func printValorByRef(valor *string) {

	fmt.Printf("Valor por referência = %v\r\n", *valor)
}

func main() {

	fmt.Printf("FUNÇÕES!!!!!!!!!!!\r\n")

	fmt.Printf("Soma 1+1 = %v\r\n", soma(1, 1))

	b, a := troca("A", "B")
	fmt.Printf("troca A, B = %v, %v\r\n", b, a)

	resu, rest := divide(5, 2)
	fmt.Printf("A divisão de 5 po 2 é %v\r\n", resu)
	fmt.Printf("o resto da divisão de 5 por 2 é = %v\r\n", rest)

	ola := func(v string) string {
		return "Olá " + v + "!\r\n"
	}

	executaFuncao(ola, "Edson")

	valor := "Esse valor não vai ser copiado, si estamos passando o ponteiro"

	printValorByRef(&valor)
}
