package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	text = strings.ToLower(text)
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es Palindromo")
	} else {
		fmt.Println("No es Palindromo")
	}
}

func main() {
	slice := []string{"Hola", "que", "hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	//sin mostrar el indice
	for _, valor := range slice {
		fmt.Println(valor)
	}

	isPalindromo("amor a roma")
	isPalindromo("Amor a roma")
}
