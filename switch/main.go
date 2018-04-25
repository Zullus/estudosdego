package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {

	fmt.Print("UNIX box?\n\r")

	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "openbsd":
		fallthrough
	case "plan9":
		fmt.Printf("YES!\r\n")
	case "linux":
		fmt.Printf("almost...\r\n")
	default:
		fmt.Printf("not at all...\r\n")
	}

	fmt.Print("Checando números de 1 a 10\r\n")

	fmt.Print("Digite um número: ")

	var inserido string

	fmt.Scanln(&inserido)

	switch numero, _ := strconv.Atoi(inserido); numero {

	case 1, 3, 5, 7:

		fmt.Printf("%v é primo! \n\r", numero)

		fallthrough

	case 9:

		fmt.Printf("%v é ímpar! \n\r", numero)

		fmt.Printf("O resto da divisão de %v por 2 é %v\n\r", numero, numero%2)

	case 4, 6, 8, 10:

		fmt.Printf("%v é par! \n\r", numero)

	case 2:

		fmt.Printf("%v é uma caso a parte!\r\n", numero)

		fmt.Printf("O número %v é o único número par quer é primo!!!\r\n", numero)

	default:

		fmt.Printf("%v não está entre 1 e 10!\r\n", numero)

	}
}
