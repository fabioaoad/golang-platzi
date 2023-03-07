package main

import "fmt"

func main() {
	//Declaración de constantes - no cambia en el tiempo
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("Hola mundo")
	fmt.Println(pi, pi2)
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// 3 formas de declaración de variables enteras - := la var es creada en ese momento
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int     //0
	var b float64 //0
	var c string  // vacio
	var d bool    // false

	fmt.Println(a, b, c, d)

	// Área de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Área cuadrado: ", areaCuadrado)

	// Operadores aritméticos
	x := 10
	y := 50

	// Suma
	result := x + y

	fmt.Println("Suma: ", result)

	// Resta
	result = y - x
	fmt.Println("Resta: ", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación: ", result)

	// División
	result = y / x
	fmt.Println("División: ", result)

	// Modulo o Residuo
	result = y % x
	fmt.Println("Modulo: ", result)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Incremental: ", x)

	// Retos
	// Rectángulo
	baseRectangulo := 20
	alturaRectangulo := 10

	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("El Area del Rectángulo es :", areaRectangulo)

	// Circulo : AreaCirculo = pi por radio al cudrado
	const PI float64 = 3.14 // Constant
	var radioCirculo float64 = 10

	areaCirculo := PI * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

	// Trapecio
	var baseUno float64 = 6
	var baseDos float64 = 15
	var alturaTrapecio float64 = 25

	areaTrapecio := ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("El Area del Trapecio es :", areaTrapecio)
}
