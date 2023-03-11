package main

import "fmt"

func main() {

	// for condicional
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	// for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// for forever
	//counterForever := 0
	//for {
	//	fmt.Println(counterForever)
	//	counterForever++
	//	}
	fmt.Printf("\n")

	// reto: iniciar en 10 e ir imprimiendo a medida que se decrementa el indice
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
	fmt.Println("///")
	fmt.Printf("2, %T", 2)
	fmt.Println("///")

	a := 10
	b := &a
	c := *b
	println(a, b, c)
}
