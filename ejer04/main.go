package main

import (
	"fmt"
	"log"
	"strconv"
)

func esPar(num int) bool {
	return num%2 == 0
}

func esUsuarioValido(userName, pass string) bool {
	return userName == "fabioaoad" && pass == "18-12-2022"
}

func main() {

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	if valor1 == 1 && valor2 == 5 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("No es verdad")
	}

	// OR
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad")
	}

	// Covertir texto a número
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value", value)

	// Saber si es par
	if esPar(58) {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}

	// Validar pass con usuario
	if esUsuarioValido("fabioaoad", "18-12-2022") {
		fmt.Println("Usuario y password correctos")
	} else {
		fmt.Println("Usuario y password incorrectos")
	}

}
