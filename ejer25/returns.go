package main

import "fmt"

// values es un slice de enteros
func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

// names es un slice de string
func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// defino funcion que NO tiene retornos con nombre
func getValuesSinRetornoDeNombre(x int) (int, int, int) {
	return 2 * x, 3 * x, 4 * x
}

// defino funcion que tiene retornos con nombre
func getValuesConRetornoDeNombre(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3, 4))

	printNames("Fabio", "Martin", "Aoad", "FMA")

	fmt.Println(getValuesSinRetornoDeNombre(2))
	fmt.Println(getValuesConRetornoDeNombre(2))
}
