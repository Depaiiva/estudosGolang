package main

import (
	"fmt"
)

func main() {
	// 1° tipo de declaração de variáveis: declaração explicita
	var variavel1 string = "variavel 1"
	fmt.Println("imprimindo: " + variavel1)

	// 2° tipo de declaração de variáveis: declaração implicita/inferência de tipo
	variavel2 := "variavel 2"
	fmt.Println("imprimindo: " + variavel2)

	// 3° declarar mais de 1 variável
	// declarando explicitamente
	var (
		variavel3 string = "variavel 3"
		variavel4 string = "variavel 4"
	)

	fmt.Println("imprimindo: " + variavel3 + " e " + variavel4)

	// declarando com inferência de tipos
	variavel5, variavel6 := "variavel 5", "variavel 6"
	fmt.Println("imprimindo: " + variavel5 + " e " + variavel6)

	// declarando constante
	const constante1 string = "constante 1"
	fmt.Println("imprimindo: " + constante1)

	// invertendo valores de variaveis sem precisar criar um variavel temporaria
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println("imprimindo: " + variavel5 + " e " + variavel6)

}
