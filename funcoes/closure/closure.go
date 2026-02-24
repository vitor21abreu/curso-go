package main

import "fmt"

func closure() func() { // tem seu propio escopo e tem ideia de onde ela ta
	// tras com sigo toda as informaçoes
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX()
}
