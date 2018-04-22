package main

import (
	"fmt"
)

var (
	nome                   = "Edson Santos"
	idade                  = 44
	url, path, query, page = "xsinformatica.com.br", "/contato", "&lang=go", 1
)

const pi = 3.141592

func main() {

	var x, y = 1, 2

	var a string
	var s string

	a = "texto 1"

	b := "texto 2"

	ola := func() {
		fmt.Printf("Olá da funcão anônima!\r\n")

	}

	fmt.Printf("a tipo: %T\n\r", a)
	fmt.Printf("b tipo: %T\r\n", b)
	fmt.Printf("╥ tipo: %T\r\n", pi)
	fmt.Printf("ola tio: %T\r\n", ola)

	fmt.Printf("O valor de a = %v\r\n", a)
	fmt.Printf("O valor de b = %v\r\n", b)
	fmt.Printf("O valor de ╥ = %v\n\r", pi)

	fmt.Printf("O valor de x = %v\n\r", x)
	fmt.Printf("O valor de y = %v\r\n", y)

	fmt.Printf("O valor de s = %v\r\n", s)

	ola()
}
