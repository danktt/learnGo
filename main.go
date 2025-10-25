package main

import (
	"fmt"

	"main.go/pacote"
) 



func main() {
	fmt.Println(pacote.Fus, pacote.Ro, pacote.Dah + "!")

	fmt.Println(soma(2, 4))
	fmt.Println(dividir(5, 2))
}



func soma( a, b  int) int {
	return a + b 
}

func dividir ( a, b int) (resultado, resto int ) {
	resultado = a / b
	resto = a % b

	return resultado, resto
}