package main

import (
	"fmt"
	"math"
)

func normalFunction() {
	fmt.Println("Hola Mundo")
}

func secondFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func dobleReturnValue(a int) (c, d int) {
	return a, a * 2
}

func areaCirculo(radio float64) float64 {
	return math.Pi * radio * radio
}
func areaRectangulo(base float64, altura float64) float64 {
	return base * altura
}

func areaTrapezoide(B float64, b float64, h float64) float64 {
	return h * (B + b) / 2
}

func main() {
	normalFunction()

	secondFunction("Hola Mundo")

	tripleArgument(1, 2, "Hola")

	value := returnValue(2)
	fmt.Println("Value:", value)

	value1, value2 := dobleReturnValue(2)
	fmt.Println("Value1 y Value2:", value1, value2)

	value3, _ := dobleReturnValue(4)
	fmt.Println("Value3:", value3)

	fmt.Printf("Circulo %.2f \n", areaCirculo(2))
	fmt.Printf("Rectangulo %.2f \n", areaRectangulo(5, 10))
	fmt.Printf("Trapezoide %.2f \n", areaTrapezoide(10, 5, 3))

}
